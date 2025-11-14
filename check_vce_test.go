package done

import (
	"testing"
)

// TestVE tests the VCE function creates a Vce instance.
// TestVE 测试 VCE 函数创建 Vce 实例。
func TestVE(t *testing.T) {
	ve := VCE(newExample())
	t.Log(ve.V)
}

// TestVe_Done tests the Done method checks error and returns value.
// TestVe_Done 测试 Done 方法检查错误并返回值。
func TestVe_Done(t *testing.T) {
	example := VCE(newExample()).Done()
	t.Log(example.S)
}

// TestVe_Good tests the Good method ensures value is not zero.
// TestVe_Good 测试 Good 方法确保值非零。
func TestVe_Good(t *testing.T) {
	VCE(newExample()).Good()
}

// TestVe_Nice tests the Nice method ensures value is not zero and returns it.
// TestVe_Nice 测试 Nice 方法确保值非零并返回值。
func TestVe_Nice(t *testing.T) {
	example := VCE(newExample()).Nice()
	t.Log(example.S)
}

// TestVe_Zero tests the Zero method ensures value is zero.
// TestVe_Zero 测试 Zero 方法确保值是零值。
func TestVe_Zero(t *testing.T) {
	VCE(newFalse()).Zero()
	VCE(newInt64(0)).Zero()
}

// newFalse returns false boolean value.
// newFalse 返回 false 布尔值。
func newFalse() (bool, error) {
	return false, nil
}

// newInt64 returns an int64 value.
// newInt64 返回一个 int64 值。
func newInt64(v int64) (int64, error) {
	return v, nil
}

// newUint64 returns an uint64 value.
// newUint64 返回一个 uint64 值。
func newUint64(v uint64) (uint64, error) {
	return v, nil
}
