package done

import (
	"testing"
)

func TestVE(t *testing.T) {
	ve := VCE(NewExample())
	t.Log(ve.V)
}

func TestVe_Done(t *testing.T) {
	example := VCE(NewExample()).Done()
	t.Log(example.S)
}

func TestVe_Good(t *testing.T) {
	VCE(NewExample()).Good()
}

func TestVe_Nice(t *testing.T) {
	example := VCE(NewExample()).Nice()
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
