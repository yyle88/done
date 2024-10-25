package done

import (
	"testing"
)

func TestVme_Nice(t *testing.T) {
	run := func() (map[int64]string, error) {
		return map[int64]string{
			1: "a",
		}, nil
	}

	res := VME(run()).Nice()
	t.Log(res)
}

func TestVme_Size(t *testing.T) {
	run := func() (map[string]float64, error) {
		return map[string]float64{
			"a": 1.0,
			"b": 2.0,
			"c": 3.0,
		}, nil
	}

	VME(run()).Size(3)
}
