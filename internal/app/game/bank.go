package game

import (
	"fmt"
	"sync"
)

const maxWallet = 1000

type bank struct {
	mu         sync.Mutex
	balance    int
	nextCoin   int
	coinIssued []bool
}

func newBank(totalCash int) *bank {
	return &bank{
		coinIssued: make([]bool, totalCash),
	}
}

func (b *bank) getBalance() (balance int, wallet []int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	walletSize := maxWallet
	if b.balance < walletSize {
		walletSize = b.balance
	}
	wallet = make([]int, walletSize)

	// Search for issued coins by looking back from last issued coin
	// because there is higher chance last issued coins wasn't spent yet.
	next := b.nextCoin - 1
	iter := newRR(next, false, len(b.coinIssued))
	for i := range wallet {
		for !b.coinIssued[next] {
			next = iter.next()
		}
		wallet[i] = next
		next = iter.next()
	}

	return b.balance, wallet
}

func (b *bank) earn(amount int) (wallet []int, _ error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if !(amount >= 1 && amount+b.balance > b.balance && amount+b.balance <= len(b.coinIssued)) {
		return nil, fmt.Errorf("%w: %d (balance=%d, overall coins=%d)", ErrWrongAmount, amount, b.balance, len(b.coinIssued))
	}

	wallet = make([]int, amount)
	iter := newRR(b.nextCoin, true, len(b.coinIssued))
	for i := range wallet {
		for b.coinIssued[b.nextCoin] {
			b.nextCoin = iter.next()
		}
		b.coinIssued[b.nextCoin] = true
		b.balance++
		wallet[i] = b.nextCoin
		b.nextCoin = iter.next()
	}
	return wallet, nil
}

func (b *bank) spend(wallet []int) (err error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if len(wallet) > b.balance {
		return ErrNotEnoughMoney
	}
	for i := range wallet {
		coin := wallet[i]
		switch {
		case coin < 0 || coin >= len(b.coinIssued):
			err = ErrCoinNotExists
		case !b.coinIssued[coin]:
			err = fmt.Errorf("%w: %d", ErrCoinNotIssued, coin)
		}
		if err != nil {
			for j := 0; j < i; j++ {
				b.coinIssued[wallet[j]] = true
			}
			return err
		}
		b.coinIssued[coin] = false
	}
	b.balance -= len(wallet)
	return nil
}
