package done

import (
	"testing"
)

func TestVE(t *testing.T) {
	ve := VCE(newExample())
	t.Log(ve.V)
}

func TestVe_Done(t *testing.T) {
	example := VCE(newExample()).Done()
	t.Log(example.S)
}

func TestVe_Good(t *testing.T) {
	VCE(newExample()).Good()
}

func TestVe_Nice(t *testing.T) {
	example := VCE(newExample()).Nice()
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
