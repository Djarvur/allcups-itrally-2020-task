package game

import (
	"fmt"
	"sync"
)

type licenses struct {
	maxActive      int
	mu             sync.Mutex
	nextID         int
	activeLicenses map[int]*License
}

func newLicenses(maxActive int) *licenses {
	return &licenses{
		maxActive:      maxActive,
		activeLicenses: make(map[int]*License, maxActive),
	}
}

func (l *licenses) active() []License {
	l.mu.Lock()
	defer l.mu.Unlock()

	active := make([]License, 0, len(l.activeLicenses))
	for id := range l.activeLicenses {
		active = append(active, *l.activeLicenses[id])
	}
	return active
}

func (l *licenses) issue(digAllowed int) (*License, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if digAllowed < 1 || digAllowed > maxDigAllowed {
		panic(fmt.Sprintf("digAllowed=%d must be between 1 and %d", digAllowed, maxDigAllowed))
	}
	if len(l.activeLicenses) >= l.maxActive {
		return nil, ErrActiveLicenseLimit
	}

	license := &License{
		ID:         l.nextID,
		DigAllowed: digAllowed,
	}
	l.activeLicenses[l.nextID] = license
	l.nextID++
	return license, nil
}

func (l *licenses) use(id int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	license, ok := l.activeLicenses[id]
	if !ok {
		return ErrNoSuchLicense
	}
	license.DigUsed++
	if license.DigUsed >= license.DigAllowed {
		delete(l.activeLicenses, id)
	}
	return nil
}
