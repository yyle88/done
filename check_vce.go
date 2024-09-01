package done

import "github.com/pkg/errors"

type Vce[V comparable] struct {
	V V
	E error
	*Ve[V]
}

// VCE accept two params. one is comparable and one is error interface
func VCE[V comparable](val V, err error) *Vce[V] {
	return &Vce[V]{
		V:  val,
		E:  err,
		Ve: VE[V](val, err),
	}
}

func (a *Vce[V]) Sure() V {
	return Sure(a.Done())
}

func (a *Vce[V]) Nice() V {
	return Nice(a.Done())
}

func (a *Vce[V]) Good() {
	Good(a.Done())
}

func (a *Vce[V]) Fine() {
	Fine(a.Done())
}

func (a *Vce[V]) Safe() {
	Safe(a.Done())
}

func (a *Vce[V]) Zero() {
	Zero(a.Done())
}

func (a *Vce[V]) None() {
	None(a.Done())
}

func (a *Vce[V]) Same(v V) {
	if x := a.Done(); x != v {
		panic(errors.New("SHOULD BE SAME BUT IS DIFFERENT"))
	}
}

func (a *Vce[V]) Diff(v V) {
	if x := a.Done(); x == v {
		panic(errors.New("SHOULD BE DIFFERENT BUT IS SAME"))
	}
}

func (a *Vce[V]) Is(v V) {
	if x := a.Done(); x != v {
		panic(errors.New("SHOULD BE SAME BUT IS DIFFERENT"))
	}
}

func (a *Vce[V]) Equals(v V) {
	if x := a.Done(); x != v {
		panic(errors.New("SHOULD BE SAME BUT IS DIFFERENT"))
	}
}

// Different 这个单词这么长真是万恶啊
func (a *Vce[V]) Different(v V) {
	if x := a.Done(); x == v {
		panic(errors.New("SHOULD BE DIFFERENT BUT IS SAME"))
	}
}
