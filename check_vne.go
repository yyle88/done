package done

import "github.com/pkg/errors"

type numType interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | uint | float32 | float64
}

type Vne[V numType] struct {
	V V
	E error
	*Vce[V]
}

func VNE[V numType](val V, err error) *Vne[V] {
	return &Vne[V]{
		V:   val,
		E:   err,
		Vce: VCE[V](val, err),
	}
}

func (a *Vne[V]) Gt(base V) V {
	if x := a.Done(); x > base {
		return x
	} else {
		panic(errors.Errorf("SHOULD BE x > BASE BUT NOT. x=%v BASE=%v", x, base))
	}
}

func (a *Vne[V]) Lt(base V) V {
	if x := a.Done(); x < base {
		return x
	} else {
		panic(errors.Errorf("SHOULD BE x < BASE BUT NOT. x=%v BASE=%v", x, base))
	}
}

func (a *Vne[V]) Gte(base V) V {
	if x := a.Done(); x >= base {
		return x
	} else {
		panic(errors.Errorf("SHOULD BE x >= BASE BUT NOT. x=%v BASE=%v", x, base))
	}
}

func (a *Vne[V]) Lte(base V) V {
	if x := a.Done(); x <= base {
		return x
	} else {
		panic(errors.Errorf("SHOULD BE x <= BASE BUT NOT. x=%v BASE=%v", x, base))
	}
}
