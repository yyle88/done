package done

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done/internal/tests"
)

func TestGood(t *testing.T) {
	Good(&Example{})
	Good(Example{S: "xyz"})
}

func TestNice(t *testing.T) {
	a := Nice(Example{S: "abc"})
	require.Equal(t, "abc", a.S)
	p := Nice(&Example{S: "uvw"})
	require.Equal(t, "uvw", p.S)
}

func TestZero(t *testing.T) {
	Zero((*Example)(nil))
	Zero(Example{})
}

type Example struct {
	S string
}

func newExample() (*Example, error) {
	return &Example{S: "xyz"}, nil
}

func TestNull(t *testing.T) {
	var example *Example
	Null(example)

	tests.ExpectPanic(t, func() {
		Null(&Example{S: "abc"})
	})
}

func TestFull(t *testing.T) {
	Full(&Example{S: "abc"})

	tests.ExpectPanic(t, func() {
		var example *Example
		Full(example)
	})
}
