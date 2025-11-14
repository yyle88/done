package done

import "testing"

// TestVBE_OK tests the OK method ensures boolean value is true.
// TestVBE_OK 测试 OK 方法确保布尔值是 true。
func TestVBE_OK(t *testing.T) {
	run := func() (bool, error) {
		return true, nil
	}

	VBE(run()).OK()
}

// TestVBE_NO tests the NO method ensures boolean value is false.
// TestVBE_NO 测试 NO 方法确保布尔值是 false。
func TestVBE_NO(t *testing.T) {
	run := func() (bool, error) {
		return false, nil
	}

	VBE(run()).NO()
}
