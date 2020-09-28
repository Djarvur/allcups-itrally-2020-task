package resource

import (
	"context"
	"fmt"
	"time"
)

type Ctx = context.Context

type CPU struct {
	freq time.Duration
	tick chan struct{}
}

func NewCPU(hz int) *CPU {
	const compensate = 2 * time.Millisecond // Compensate slow consumer for up to 0.002s.
	if hz < 1 {
		panic(fmt.Sprintf("hz must be a positive number: %d", hz))
	}
	return &CPU{
		freq: time.Second / time.Duration(hz),
		tick: make(chan struct{}, hz/int(time.Second/compensate)),
	}
}

func (c *CPU) Provide(ctx Ctx) error {
	prev := time.Now().Round(c.freq)
	tickc := time.Tick(time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			return nil
		case now := <-tickc:
			now = now.Round(c.freq)
			for i := 0; i < int(now.Sub(prev)/c.freq); i++ {
				select {
				case c.tick <- struct{}{}:
				default:
				}
			}
			prev = now
		}
	}
}

func (c *CPU) Consume(ctx Ctx, t time.Duration) error {
	for ticks := int(t.Round(c.freq) / c.freq); ticks > 0; ticks-- {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c.tick:
		}
	}
	return nil
}
