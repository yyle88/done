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

func TestV3(t *testing.T) {
	run := func() (string, uint64, rune, error) {
		return "a", 2, 'x', nil
	}

	v1, v2, v3 := done.V3(run())
	require.Equal(t, "a", v1)
	require.Equal(t, uint64(2), v2)
	require.Equal(t, 'x', v3)
}

func TestV4(t *testing.T) {
	run := func() (string, uint64, rune, int32, error) {
		return "a", 2, 'x', 42, nil
	}

	v1, v2, v3, v4 := done.V4(run())
	require.Equal(t, "a", v1)
	require.Equal(t, uint64(2), v2)
	require.Equal(t, 'x', v3)
	require.Equal(t, int32(42), v4)
}

func TestV5(t *testing.T) {
	run := func() (string, uint64, rune, int32, float64, error) {
		return "a", 2, 'x', 42, 3.14, nil
	}

	v1, v2, v3, v4, v5 := done.V5(run())
	require.Equal(t, "a", v1)
	require.Equal(t, uint64(2), v2)
	require.Equal(t, 'x', v3)
	require.Equal(t, int32(42), v4)
	require.Equal(t, 3.14, v5)
}

func TestV6(t *testing.T) {
	run := func() (string, uint64, rune, int32, float64, bool, error) {
		return "a", 2, 'x', 42, 3.14, true, nil
	}

	v1, v2, v3, v4, v5, v6 := done.V6(run())
	require.Equal(t, "a", v1)
	require.Equal(t, uint64(2), v2)
	require.Equal(t, 'x', v3)
	require.Equal(t, int32(42), v4)
	require.Equal(t, 3.14, v5)
	require.Equal(t, true, v6)
}

func TestV7(t *testing.T) {
	run := func() (string, uint64, rune, int32, float64, bool, byte, error) {
		return "a", 2, 'x', 42, 3.14, true, byte(255), nil
	}

	v1, v2, v3, v4, v5, v6, v7 := done.V7(run())
	require.Equal(t, "a", v1)
	require.Equal(t, uint64(2), v2)
	require.Equal(t, 'x', v3)
	require.Equal(t, int32(42), v4)
	require.Equal(t, 3.14, v5)
	require.Equal(t, true, v6)
	require.Equal(t, byte(255), v7)
}

func TestV8(t *testing.T) {
	run := func() (string, uint64, rune, int32, float64, bool, byte, int64, error) {
		return "a", 2, 'x', 42, 3.14, true, byte(255), int64(123456789), nil
	}

	v1, v2, v3, v4, v5, v6, v7, v8 := done.V8(run())
	require.Equal(t, "a", v1)
	require.Equal(t, uint64(2), v2)
	require.Equal(t, 'x', v3)
	require.Equal(t, int32(42), v4)
	require.Equal(t, 3.14, v5)
	require.Equal(t, true, v6)
	require.Equal(t, byte(255), v7)
	require.Equal(t, int64(123456789), v8)
}

func TestV9(t *testing.T) {
	run := func() (string, uint64, rune, int32, float64, bool, byte, int64, uint32, error) {
		return "a", 2, 'x', 42, 3.14, true, byte(255), int64(123456789), uint32(987654321), nil
	}

	v1, v2, v3, v4, v5, v6, v7, v8, v9 := done.V9(run())
	require.Equal(t, "a", v1)
	require.Equal(t, uint64(2), v2)
	require.Equal(t, 'x', v3)
	require.Equal(t, int32(42), v4)
	require.Equal(t, 3.14, v5)
	require.Equal(t, true, v6)
	require.Equal(t, byte(255), v7)
	require.Equal(t, int64(123456789), v8)
	require.Equal(t, uint32(987654321), v9)
}
