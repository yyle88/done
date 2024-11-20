package done_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
)

func TestP0(t *testing.T) {
	done.P0(error(nil))
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

func TestP3(t *testing.T) {
	run := func() (*string, *int, *bool, error) {
		v1 := "test"
		v2 := 100
		v3 := true
		return &v1, &v2, &v3, nil
	}

	v1, v2, v3 := done.P3(run())
	require.Equal(t, "test", *v1)
	require.Equal(t, 100, *v2)
	require.Equal(t, true, *v3)
}

func TestP4(t *testing.T) {
	run := func() (*int, *float64, *bool, *string, error) {
		v1 := 42
		v2 := 3.14
		v3 := false
		v4 := "world"
		return &v1, &v2, &v3, &v4, nil
	}

	v1, v2, v3, v4 := done.P4(run())
	require.Equal(t, 42, *v1)
	require.Equal(t, 3.14, *v2)
	require.Equal(t, false, *v3)
	require.Equal(t, "world", *v4)
}

func TestP5(t *testing.T) {
	run := func() (*string, *int, *bool, *float64, *rune, error) {
		v1 := "test"
		v2 := 123
		v3 := true
		v4 := 2.718
		v5 := 'a'
		return &v1, &v2, &v3, &v4, &v5, nil
	}

	v1, v2, v3, v4, v5 := done.P5(run())
	require.Equal(t, "test", *v1)
	require.Equal(t, 123, *v2)
	require.Equal(t, true, *v3)
	require.Equal(t, 2.718, *v4)
	require.Equal(t, 'a', *v5)
}

func TestP6(t *testing.T) {
	run := func() (*float64, *int, *string, *bool, *rune, *int64, error) {
		v1 := 3.14
		v2 := 42
		v3 := "hello"
		v4 := true
		v5 := 'x'
		v6 := int64(100000)
		return &v1, &v2, &v3, &v4, &v5, &v6, nil
	}

	v1, v2, v3, v4, v5, v6 := done.P6(run())
	require.Equal(t, 3.14, *v1)
	require.Equal(t, 42, *v2)
	require.Equal(t, "hello", *v3)
	require.Equal(t, true, *v4)
	require.Equal(t, 'x', *v5)
	require.Equal(t, int64(100000), *v6)
}

func TestP7(t *testing.T) {
	run := func() (*bool, *string, *int, *float64, *rune, *int64, *uint32, error) {
		v1 := true
		v2 := "world"
		v3 := 50
		v4 := 1.618
		v5 := 'y'
		v6 := int64(200000)
		v7 := uint32(300)
		return &v1, &v2, &v3, &v4, &v5, &v6, &v7, nil
	}

	v1, v2, v3, v4, v5, v6, v7 := done.P7(run())
	require.Equal(t, true, *v1)
	require.Equal(t, "world", *v2)
	require.Equal(t, 50, *v3)
	require.Equal(t, 1.618, *v4)
	require.Equal(t, 'y', *v5)
	require.Equal(t, int64(200000), *v6)
	require.Equal(t, uint32(300), *v7)
}

func TestP8(t *testing.T) {
	run := func() (*string, *int, *bool, *float64, *rune, *int64, *uint32, *uint8, error) {
		v1 := "abc"
		v2 := 999
		v3 := false
		v4 := 0.577
		v5 := 'z'
		v6 := int64(500000)
		v7 := uint32(400)
		v8 := uint8(255)
		return &v1, &v2, &v3, &v4, &v5, &v6, &v7, &v8, nil
	}

	v1, v2, v3, v4, v5, v6, v7, v8 := done.P8(run())
	require.Equal(t, "abc", *v1)
	require.Equal(t, 999, *v2)
	require.Equal(t, false, *v3)
	require.Equal(t, 0.577, *v4)
	require.Equal(t, 'z', *v5)
	require.Equal(t, int64(500000), *v6)
	require.Equal(t, uint32(400), *v7)
	require.Equal(t, uint8(255), *v8)
}

func TestP9(t *testing.T) {
	run := func() (*string, *int, *bool, *float64, *rune, *int64, *uint32, *uint8, *uint64, error) {
		v1 := "final"
		v2 := 888
		v3 := true
		v4 := 2.718
		v5 := 'p'
		v6 := int64(900000)
		v7 := uint32(500)
		v8 := uint8(100)
		v9 := uint64(1000000)
		return &v1, &v2, &v3, &v4, &v5, &v6, &v7, &v8, &v9, nil
	}

	v1, v2, v3, v4, v5, v6, v7, v8, v9 := done.P9(run())
	require.Equal(t, "final", *v1)
	require.Equal(t, 888, *v2)
	require.Equal(t, true, *v3)
	require.Equal(t, 2.718, *v4)
	require.Equal(t, 'p', *v5)
	require.Equal(t, int64(900000), *v6)
	require.Equal(t, uint32(500), *v7)
	require.Equal(t, uint8(100), *v8)
	require.Equal(t, uint64(1000000), *v9)
}
