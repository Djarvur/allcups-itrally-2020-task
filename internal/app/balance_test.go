package app_test

import (
	"testing"

	"github.com/powerman/check"
)

func TestBalance(tt *testing.T) {
	t := check.T(tt)
	a, _ := testNew(t)

	t.TODO().NotPanic(func() { a.Balance(ctx) })
}
