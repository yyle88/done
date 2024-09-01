package done

import "testing"

func TestVBE_OK(t *testing.T) {
	VBE(newExample1x(true)).OK()
}

func TestVBE_NO(t *testing.T) {
	VBE(newExample1x(false)).NO()
}
