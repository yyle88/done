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

func (a *Vne[V]) Gt(v V) V {
	if x := a.Done(); !(x > v) {
		panic(errors.Errorf("SHOULD BE x > v BUT NOT. x=%v v=%v", x, v))
	}
	return v
}

func (a *Vne[V]) Lt(v V) V {
	if x := a.Done(); !(x < v) {
		panic(errors.Errorf("SHOULD BE x < v BUT NOT. x=%v v=%v", x, v))
	}
	return v
}

func (a *Vne[V]) Gte(v V) V {
	if x := a.Done(); !(x >= v) {
		panic(errors.Errorf("SHOULD BE x >= v BUT NOT. x=%v v=%v", x, v))
	}
	return v
}

func (a *Vne[V]) Lte(v V) V {
	if x := a.Done(); !(x <= v) {
		panic(errors.Errorf("SHOULD BE x <= v BUT NOT. x=%v v=%v", x, v))
	}
	return v
}
