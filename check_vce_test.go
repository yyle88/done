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
	VCE(newInt64()).Zero()
}

func newFalse() (bool, error) {
	return false, nil
}

func newInt64() (int64, error) {
	return 0, nil
}
