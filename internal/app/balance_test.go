package app_test

import (
	"testing"

	"github.com/powerman/check"
)

func TestBalance(tt *testing.T) {
	t := check.T(tt)
	cleanup, a, _ := testNew(t)
	defer cleanup()

	t.TODO().NotPanic(func() { a.Balance(ctx) })
}
