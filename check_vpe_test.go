package done

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestVpe_Sure tests the Sure method ensures pointer is not nil and returns it.
// TestVpe_Sure 测试 Sure 方法确保指针非 nil 并返回指针。
func TestVpe_Sure(t *testing.T) {
	run := func() (*int32, error) {
		num := int32(100)
		return &num, nil
	}
	p := VPE(run()).Sure()
	t.Log(*p)
	require.Equal(t, int32(100), *p)
}

// TestVpe_Nice tests the Nice method ensures pointer is not nil and returns it.
// TestVpe_Nice 测试 Nice 方法确保指针非 nil 并返回指针。
func TestVpe_Nice(t *testing.T) {
	type exampleType struct {
		num int64
	}

	run := func() (*exampleType, error) {
		return &exampleType{num: 200}, nil
	}
	p := VPE(run()).Sure()
	t.Log(p)
	require.Equal(t, int64(200), p.num)
}

// TestVpe_None tests the None method ensures pointer is nil.
// TestVpe_None 测试 None 方法确保指针是 nil。
func TestVpe_None(t *testing.T) {
	run := func() (*int64, error) {
		return nil, nil
	}
	VPE(run()).None()
}

// TestVpe_Full tests the Full method ensures pointer is not nil and returns it.
// TestVpe_Full 测试 Full 方法确保指针非 nil 并返回指针。
func TestVpe_Full(t *testing.T) {
	type exampleType struct {
		Value int64
	}
	run := func() (*exampleType, error) {
		return &exampleType{Value: 666}, nil
	}
	res := VPE(run()).Full()
	require.Equal(t, int64(666), res.Value)
}
