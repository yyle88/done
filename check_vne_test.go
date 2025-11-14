package done

import "testing"

// TestVne_Zero tests the Zero and Nice methods for numeric values.
// TestVne_Zero 测试数字值的 Zero 和 Nice 方法。
func TestVne_Zero(t *testing.T) {
	VNE(newInt64(0)).Zero()
	VNE(newInt64(1)).Nice()
}

// TestVne_Same tests the Same method ensures value equals given value.
// TestVne_Same 测试 Same 方法确保值等于给定值。
func TestVne_Same(t *testing.T) {
	VNE(newInt64(0)).Same(0)
	VNE(newInt64(1)).Same(1)
}

// TestVne_Is tests the Is method ensures value equals given value.
// TestVne_Is 测试 Is 方法确保值等于给定值。
func TestVne_Is(t *testing.T) {
	VNE(newInt64(200)).Is(200)
	VNE(newInt64(404)).Is(404)
}

// TestVne_Gt tests comparison methods (Gt, Lt, Gte, Lte) for int64.
// TestVne_Gt 测试 int64 的比较方法（Gt、Lt、Gte、Lte）。
func TestVne_Gt(t *testing.T) {
	t.Log(VNE(newInt64(3)).Gt(2))
	t.Log(VNE(newInt64(2)).Lt(3))
	t.Log(VNE(newInt64(3)).Gte(3))
	t.Log(VNE(newInt64(2)).Lte(2))
}

// TestVne_Gt_Uint64 tests comparison methods (Gt, Lt, Gte, Lte) for uint64.
// TestVne_Gt_Uint64 测试 uint64 的比较方法（Gt、Lt、Gte、Lte）。
func TestVne_Gt_Uint64(t *testing.T) {
	// Compile-time check prevents passing -1 to uint64 // 编译期检查阻止传递 -1 给 uint64
	t.Log(VNE(newUint64(3)).Gt(2))
	t.Log(VNE(newUint64(2)).Lt(3))
	t.Log(VNE(newUint64(3)).Gte(3))
	t.Log(VNE(newUint64(2)).Lte(2))
}
