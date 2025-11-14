package done

import "testing"

// TestVAE_Length tests the Length method ensures slice has given length.
// TestVAE_Length 测试 Length 方法确保切片长度等于给定值。
func TestVAE_Length(t *testing.T) {
	run := func() ([]int, error) {
		return []int{1, 2, 3, 4, 5}, nil
	}

	VAE(run()).Length(5)
}

// TestVAE_Len tests the Len method ensures slice has given length.
// TestVAE_Len 测试 Len 方法确保切片长度等于给定值。
func TestVAE_Len(t *testing.T) {
	run := func() ([]string, error) {
		return []string{"a", "b", "c"}, nil
	}

	VAE(run()).Len(3)
}
