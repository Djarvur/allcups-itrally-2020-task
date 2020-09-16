package game

import (
	"fmt"
	"sync"
)

type licenses struct {
	maxActive int
	mu        sync.Mutex
	nextID    int
	licenses  map[int]*License
	isActive  map[int]bool
}

func newLicenses(maxActive int) *licenses {
	return &licenses{
		maxActive: maxActive,
		licenses:  make(map[int]*License, maxActive),
		isActive:  make(map[int]bool, maxActive),
	}
}

func (l *licenses) active() []License {
	l.mu.Lock()
	defer l.mu.Unlock()

	active := make([]License, 0, len(l.licenses))
	for id := range l.licenses {
		if l.isActive[id] {
			active = append(active, *l.licenses[id])
		}
	}
	return active
}

func (l *licenses) beginIssue(digAllowed int) (*License, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if digAllowed < 1 || digAllowed > maxDigAllowed {
		panic(fmt.Sprintf("digAllowed=%d must be between 1 and %d", digAllowed, maxDigAllowed))
	}
	if len(l.licenses) >= l.maxActive {
		return nil, ErrActiveLicenseLimit
	}

	license := &License{
		ID:         l.nextID,
		DigAllowed: digAllowed,
	}
	l.licenses[l.nextID] = license
	l.nextID++
	return license, nil
}

func (l *licenses) mustBegunIssue(id int) {
	if _, ok := l.isActive[id]; ok {
		panic("never here")
	} else if l.licenses[id] == nil {
		panic("never here")
	}
}

func (l *licenses) commitIssue(id int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.mustBegunIssue(id)
	l.isActive[id] = true
}

func (l *licenses) rollbackIssue(id int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.mustBegunIssue(id)
	delete(l.licenses, id)
}

func (l *licenses) use(id int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if _, ok := l.isActive[id]; !ok {
		return ErrNoSuchLicense
	}
	l.licenses[id].DigUsed++
	if l.licenses[id].DigUsed >= l.licenses[id].DigAllowed {
		delete(l.licenses, id)
		delete(l.isActive, id)
	}
	return nil
}
