package done

import (
	"testing"
)

func TestVE(t *testing.T) {
	ve := VCE(newExample2x())
	t.Log(ve.V)
}

func TestVe_Done(t *testing.T) {
	example := VCE(newExample2x()).Done()
	t.Log(example.S)
}

func TestVe_Good(t *testing.T) {
	VCE(newExample2x()).Good()
}

func TestVe_Nice(t *testing.T) {
	example := VCE(newExample2x()).Nice()
	t.Log(example.S)
}

func TestVe_Zero(t *testing.T) {
	VCE(newFalse()).Zero()
	VCE(newInt64(0)).Zero()
}

func newFalse() (bool, error) {
	return false, nil
}

func newInt64(v int64) (int64, error) {
	return v, nil
}

func newUint64(v uint64) (uint64, error) {
	return v, nil
}
