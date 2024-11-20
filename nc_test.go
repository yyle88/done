package done_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
)

func TestC0(t *testing.T) {
	done.C0(error(nil))
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

func TestC3(t *testing.T) {
	run := func() (bool, byte, rune, error) {
		return true, 'A', '中', nil
	}

	v1, v2, v3 := done.C3(run())
	require.True(t, v1)
	require.Equal(t, byte('A'), v2)
	require.Equal(t, '中', v3)
}

func TestC4(t *testing.T) {
	run := func() (string, int, uint, float32, error) {
		return "world", -7, 7, 1.23, nil
	}

	v1, v2, v3, v4 := done.C4(run())
	require.Equal(t, "world", v1)
	require.Equal(t, -7, v2)
	require.Equal(t, uint(7), v3)
	require.Equal(t, float32(1.23), v4)
}

func TestC5(t *testing.T) {
	run := func() (uint8, uint16, uint32, uint64, string, error) {
		return 8, 16, 32, 64, "test", nil
	}

	v1, v2, v3, v4, v5 := done.C5(run())
	require.Equal(t, uint8(8), v1)
	require.Equal(t, uint16(16), v2)
	require.Equal(t, uint32(32), v3)
	require.Equal(t, uint64(64), v4)
	require.Equal(t, "test", v5)
}

func TestC6(t *testing.T) {
	run := func() (float64, complex64, complex128, bool, string, int, error) {
		return 6.28, complex(1, 2), complex(3, 4), true, "check", -1, nil
	}

	v1, v2, v3, v4, v5, v6 := done.C6(run())
	require.Equal(t, 6.28, v1)
	require.Equal(t, complex64(complex(1, 2)), v2)
	require.Equal(t, complex(3, 4), v3)
	require.True(t, v4)
	require.Equal(t, "check", v5)
	require.Equal(t, -1, v6)
}

func TestC7(t *testing.T) {
	run := func() (uint, uint8, uint16, uint32, uint64, int, int8, error) {
		return 1, 2, 3, 4, 5, 6, 7, nil
	}

	v1, v2, v3, v4, v5, v6, v7 := done.C7(run())
	require.Equal(t, uint(1), v1)
	require.Equal(t, uint8(2), v2)
	require.Equal(t, uint16(3), v3)
	require.Equal(t, uint32(4), v4)
	require.Equal(t, uint64(5), v5)
	require.Equal(t, 6, v6)
	require.Equal(t, int8(7), v7)
}

func TestC8(t *testing.T) {
	run := func() (string, rune, byte, bool, float32, float64, int32, int64, error) {
		return "go", '中', 'B', true, 2.71, 1.618, -32, -64, nil
	}

	v1, v2, v3, v4, v5, v6, v7, v8 := done.C8(run())
	require.Equal(t, "go", v1)
	require.Equal(t, '中', v2)
	require.Equal(t, byte('B'), v3)
	require.True(t, v4)
	require.Equal(t, float32(2.71), v5)
	require.Equal(t, 1.618, v6)
	require.Equal(t, int32(-32), v7)
	require.Equal(t, int64(-64), v8)
}

func TestC9(t *testing.T) {
	run := func() (string, uint8, int16, uint32, int64, float32, float64, complex64, complex128, error) {
		return "done", 255, -16, 32, -64, 3.14, 2.71, complex(5, 6), complex(7, 8), nil
	}

	v1, v2, v3, v4, v5, v6, v7, v8, v9 := done.C9(run())
	require.Equal(t, "done", v1)
	require.Equal(t, uint8(255), v2)
	require.Equal(t, int16(-16), v3)
	require.Equal(t, uint32(32), v4)
	require.Equal(t, int64(-64), v5)
	require.Equal(t, float32(3.14), v6)
	require.Equal(t, 2.71, v7)
	require.Equal(t, complex64(complex(5, 6)), v8)
	require.Equal(t, complex(7, 8), v9)
}
