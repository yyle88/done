package done

import "testing"

func TestVne_Zero(t *testing.T) {
	VNE(newInt64(0)).Zero()
	VNE(newInt64(1)).Nice()
}

func TestVne_Same(t *testing.T) {
	VNE(newInt64(0)).Same(0)
	VNE(newInt64(1)).Same(1)
}

func TestVne_Is(t *testing.T) {
	VNE(newInt64(200)).Is(200)
	VNE(newInt64(404)).Is(404)
}

func TestVne_Gt(t *testing.T) {
	t.Log(VNE(newInt64(3)).Gt(2))
	t.Log(VNE(newInt64(2)).Lt(3))
	t.Log(VNE(newInt64(3)).Gte(3))
	t.Log(VNE(newInt64(2)).Lte(2))
}

func TestVne_Gt_Uint64(t *testing.T) {
	//t.Log(VNE(newUint64(3)).Gt(-1)) //在编译阶段即会报错，因为 -1 传不到 uint64 里面，认为这样也是符合预期的
	t.Log(VNE(newUint64(3)).Gt(2))
	t.Log(VNE(newUint64(2)).Lt(3))
	t.Log(VNE(newUint64(3)).Gte(3))
	t.Log(VNE(newUint64(2)).Lte(2))
}
