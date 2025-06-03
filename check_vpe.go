package done

type Vpe[V any] struct {
	V       *V
	E       error
	*Ve[*V] //这里最好不要继承自 vce-comparable 这个类型，因为指针比较地址相等或者不想等的意义不太大，因此不要提供这些无意义的功能
}

// VPE accept two params. one is POINTER and one is ERROR interface
func VPE[V any](val *V, err error) *Vpe[V] {
	return &Vpe[V]{
		V:  val,
		E:  err,
		Ve: VE[*V](val, err),
	}
}

func (a *Vpe[V]) Sure() *V {
	return Sure(a.Done())
}

func (a *Vpe[V]) Nice() *V {
	return Nice(a.Done())
}

func (a *Vpe[V]) Good() {
	Good(a.Done())
}

func (a *Vpe[V]) Fine() {
	Fine(a.Done())
}

func (a *Vpe[V]) Safe() {
	Safe(a.Done())
}

func (a *Vpe[V]) Zero() {
	Zero(a.Done())
}

func (a *Vpe[V]) None() {
	None(a.Done())
}

func (a *Vpe[V]) Null() {
	Null(a.Done())
}

func (a *Vpe[V]) Full() *V {
	return Full(a.Done())
}
