package done_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
)

func TestV0(t *testing.T) {
	run := func() error {
		return nil
	}

	done.V0(run())
}

func TestV1(t *testing.T) {
	run := func() (string, error) {
		return "a", nil
	}

	v1 := done.V1(run())
	require.Equal(t, "a", v1)
}

func TestV2(t *testing.T) {
	run := func() (string, uint64, error) {
		return "a", 2, nil
	}

	v1, v2 := done.V2(run())
	require.Equal(t, "a", v1)
	require.Equal(t, uint64(2), v2)
}

func TestP0(t *testing.T) {
	run := func() error {
		return nil
	}

	done.P0(run())
}

func TestP1(t *testing.T) {
	run := func() (*string, error) {
		v1 := "hello"
		return &v1, nil
	}

	p1 := done.P1(run())
	require.Equal(t, "hello", *p1)
}

func TestP2(t *testing.T) {
	run := func() (*int, *float64, error) {
		v1 := 42
		v2 := 3.14
		return &v1, &v2, nil
	}

	v1, v2 := done.P2(run())
	require.Equal(t, 42, *v1)
	require.Equal(t, 3.14, *v2)
}

func TestC0(t *testing.T) {
	run := func() error {
		return nil
	}

	done.C0(run())
}

func TestC1(t *testing.T) {
	run := func() (string, error) {
		return "hello", nil
	}

	v1 := done.C1(run())
	require.Equal(t, "hello", v1)
}

func TestC2(t *testing.T) {
	run := func() (int, float64, error) {
		return 42, 3.14, nil
	}

	v1, v2 := done.C2(run())
	require.Equal(t, 42, v1)
	require.Equal(t, 3.14, v2)
}
