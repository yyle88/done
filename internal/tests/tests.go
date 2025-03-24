package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func ExpectPanic(t *testing.T, run func()) {
	defer func() {
		if recoverFrom := recover(); recoverFrom != nil {
			t.Logf("expect panic then catch panic [%v] -> [SUCCESS]", recoverFrom)
			return
		}
		require.Fail(t, "expect panic while not panic -> [FAILURE]")
	}()

	run()
}
