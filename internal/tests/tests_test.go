package tests_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/yyle88/must/internal/tests"
)

func TestExpectPanic(t *testing.T) {
	tests.ExpectPanic(t, func() {
		panic(errors.New("expect-panic"))
	})
}
