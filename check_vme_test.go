package done

import (
	"testing"
)

// TestVme_Nice tests the Nice method ensures map is not empty and returns it.
// TestVme_Nice 测试 Nice 方法确保 map 非空并返回 map。
func TestVme_Nice(t *testing.T) {
	run := func() (map[int64]string, error) {
		return map[int64]string{
			1: "a",
		}, nil
	}

	res := VME(run()).Nice()
	t.Log(res)
}

// TestVme_Size tests the Size method ensures map has given size.
// TestVme_Size 测试 Size 方法确保 map 大小等于给定值。
func TestVme_Size(t *testing.T) {
	run := func() (map[string]float64, error) {
		return map[string]float64{
			"a": 1.0,
			"b": 2.0,
			"c": 3.0,
		}, nil
	}

	VME(run()).Size(3)
}
