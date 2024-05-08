package done

import (
	"testing"

	"github.com/stretchr/testify/require"
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

func NewExample() (*Example, error) {
	return &Example{S: "xyz"}, nil
}
