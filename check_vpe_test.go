package done

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVpe_Sure(t *testing.T) {
	run := func() (*int32, error) {
		num := int32(100)
		return &num, nil
	}
	p := VPE(run()).Sure()
	t.Log(*p)
	require.Equal(t, int32(100), *p)
}

func TestVpe_Nice(t *testing.T) {
	type exampleType struct {
		num int64
	}

	run := func() (*exampleType, error) {
		return &exampleType{num: 200}, nil
	}
	p := VPE(run()).Sure()
	t.Log(p)
	require.Equal(t, int64(200), p.num)
}

func TestVpe_None(t *testing.T) {
	run := func() (*int64, error) {
		return nil, nil
	}
	VPE(run()).None()
}

func TestVpe_Full(t *testing.T) {
	type exampleType struct {
		Value int64
	}
	run := func() (*exampleType, error) {
		return &exampleType{Value: 666}, nil
	}
	res := VPE(run()).Full()
	require.Equal(t, int64(666), res.Value)
}
