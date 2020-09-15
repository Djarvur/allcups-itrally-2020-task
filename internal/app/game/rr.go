package game

type rr struct {
	start   int // 0..length-1
	step    int // 1 or -1
	length  int // 1..
	current int // 0..length-1
}

func newRR(start int, forward bool, length int) *rr {
	if start < 0 || start >= length || length <= 0 {
		panic("never here")
	}
	step := 1
	if !forward {
		step = -1
	}
	return &rr{
		start:   start,
		step:    step,
		length:  length,
		current: start,
	}
}

// Next returns next value (moving by one from start value) until it reach
// start value again (then it panics).
func (r *rr) next() int {
	r.current += r.step
	if r.current >= r.length {
		r.current = 0
	} else if r.current < 0 {
		r.current = r.length - 1
	}
	if r.current == r.start {
		panic("never here")
	}
	return r.current
}
