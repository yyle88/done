package done

import "github.com/pkg/errors"

type Vme[K comparable, V any] struct {
	V map[K]V
	E error
	*Ve[map[K]V]
}

// VME accept two params. one is map and one is error interface
// So M means Map type.
func VME[K comparable, V any](val map[K]V, err error) *Vme[K, V] {
	return &Vme[K, V]{
		V:  val,
		E:  err,
		Ve: VE[map[K]V](val, err),
	}
}

func (a *Vme[K, V]) Sure() map[K]V {
	return a.sure1x1nice1x1some(a.Done())
}

func (a *Vme[K, V]) Nice() map[K]V {
	return a.sure1x1nice1x1some(a.Done())
}

func (a *Vme[K, V]) Some() map[K]V {
	return a.sure1x1nice1x1some(a.Done())
}

func (a *Vme[K, V]) Good() {
	a.good1x1fine1x1safe1x1have(a.Done())
}

func (a *Vme[K, V]) Fine() {
	a.good1x1fine1x1safe1x1have(a.Done())
}

func (a *Vme[K, V]) Safe() {
	a.good1x1fine1x1safe1x1have(a.Done())
}

func (a *Vme[K, V]) Have() {
	a.good1x1fine1x1safe1x1have(a.Done())
}

func (a *Vme[K, V]) Zero() {
	a.none1x1zero1x1void(a.Done())
}

func (a *Vme[K, V]) None() {
	a.none1x1zero1x1void(a.Done())
}

func (a *Vme[K, V]) Void() {
	a.none1x1zero1x1void(a.Done())
}

func (a *Vme[K, V]) Size(n int) {
	a.size1x1len1x1length(a.Done(), n)
}

func (a *Vme[K, V]) Len(n int) {
	a.size1x1len1x1length(a.Done(), n)
}

func (a *Vme[K, V]) Length(n int) {
	a.size1x1len1x1length(a.Done(), n)
}

func (a *Vme[K, V]) sure1x1nice1x1some(v map[K]V) map[K]V {
	if len(v) == 0 {
		panic(errors.New("SHOULD HAVE SOME BUT IS NONE"))
	}
	return v
}

func (a *Vme[K, V]) good1x1fine1x1safe1x1have(v map[K]V) {
	if len(v) == 0 {
		panic(errors.New("SHOULD HAVE SOME BUT IS NONE"))
	}
}

func (a *Vme[K, V]) none1x1zero1x1void(v map[K]V) {
	if len(v) != 0 {
		panic(errors.New("SHOULD BE NONE BUT HAVE SOME"))
	}
}

func (a *Vme[K, V]) size1x1len1x1length(v map[K]V, n int) {
	if len(v) != n {
		panic(errors.New("SIZE IS NOT SAME"))
	}
}
