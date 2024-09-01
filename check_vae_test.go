package done

import "testing"

func TestVAE_Length(t *testing.T) {
	VAE(newExample1x([]int{1, 2, 3, 4, 5})).Length(5)
}

func TestVAE_Len(t *testing.T) {
	VAE(newExample1x([]string{"a", "b", "c"})).Len(3)
}

func newExample1x[T any](a T) (T, error) {
	return a, nil
}
