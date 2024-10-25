package done

import "testing"

func TestVAE_Length(t *testing.T) {
	run := func() ([]int, error) {
		return []int{1, 2, 3, 4, 5}, nil
	}

	VAE(run()).Length(5)
}

func TestVAE_Len(t *testing.T) {
	run := func() ([]string, error) {
		return []string{"a", "b", "c"}, nil
	}

	VAE(run()).Len(3)
}
