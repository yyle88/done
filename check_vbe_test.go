package done

import "testing"

func TestVBE_OK(t *testing.T) {
	run := func() (bool, error) {
		return true, nil
	}

	VBE(run()).OK()
}

func TestVBE_NO(t *testing.T) {
	run := func() (bool, error) {
		return false, nil
	}

	VBE(run()).NO()
}
