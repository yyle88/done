package done_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
)

// TestV0 tests the V0 function checks error and panics if exists.
// TestV0 测试 V0 函数检查错误，存在则触发 panic。
func TestV0(t *testing.T) {
	run := func() error {
		return nil
	}

	done.V0(run())
}

// TestV1 tests the V1 function checks error and returns one value.
// TestV1 测试 V1 函数检查错误并返回一个值。
func TestV1(t *testing.T) {
	run := func() (string, error) {
		return "a", nil
	}

	v1 := done.V1(run())
	require.Equal(t, "a", v1)
}

// TestV2 tests the V2 function checks error and returns two values.
// TestV2 测试 V2 函数检查错误并返回两个值。
func TestV2(t *testing.T) {
	run := func() (string, uint64, error) {
		return "a", 2, nil
	}

	v1, v2 := done.V2(run())
	require.Equal(t, "a", v1)
	require.Equal(t, uint64(2), v2)
}

// TestP0 tests the P0 function checks error and panics if exists.
// TestP0 测试 P0 函数检查错误，存在则触发 panic。
func TestP0(t *testing.T) {
	run := func() error {
		return nil
	}

	done.P0(run())
}

// TestP1 tests the P1 function checks error, ensures pointer is not nil, and returns it.
// TestP1 测试 P1 函数检查错误，确保指针非 nil，并返回指针。
func TestP1(t *testing.T) {
	run := func() (*string, error) {
		v1 := "hello"
		return &v1, nil
	}

	p1 := done.P1(run())
	require.Equal(t, "hello", *p1)
}

// TestP2 tests the P2 function checks error, ensures pointers are not nil, and returns them.
// TestP2 测试 P2 函数检查错误，确保指针非 nil，并返回指针。
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

// TestC0 tests the C0 function checks error and panics if exists.
// TestC0 测试 C0 函数检查错误，存在则触发 panic。
func TestC0(t *testing.T) {
	run := func() error {
		return nil
	}

	done.C0(run())
}

// TestC1 tests the C1 function checks error, ensures value is not zero, and returns it.
// TestC1 测试 C1 函数检查错误，确保值非零，并返回值。
func TestC1(t *testing.T) {
	run := func() (string, error) {
		return "hello", nil
	}

	v1 := done.C1(run())
	require.Equal(t, "hello", v1)
}

// TestC2 tests the C2 function checks error, ensures values are not zero, and returns them.
// TestC2 测试 C2 函数检查错误，确保值非零，并返回值。
func TestC2(t *testing.T) {
	run := func() (int, float64, error) {
		return 42, 3.14, nil
	}

	v1, v2 := done.C2(run())
	require.Equal(t, 42, v1)
	require.Equal(t, 3.14, v2)
}
