package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func ExpectPanic(t *testing.T, run func()) {
	defer func() {
		if cause := recover(); cause != nil {
			t.Logf("expect panic then catch panic [%v] -> [SUCCESS]", cause)
			return
		}
		require.Fail(t, "expect panic while not panic -> [FAILURE]")
	}()

	run()
}
