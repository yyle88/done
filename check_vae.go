package done

import "github.com/pkg/errors"

type Vae[V any] struct {
	V []V
	E error
	*Ve[[]V]
}

// VAE accept two params. one is slice and one is error interface
// So A means Array type.
func VAE[V any](val []V, err error) *Vae[V] {
	return &Vae[V]{
		V:  val,
		E:  err,
		Ve: VE[[]V](val, err),
	}
}

func (a *Vae[V]) Sure() []V {
	return a.sure1x1nice1x1some(a.Done())
}

func (a *Vae[V]) Nice() []V {
	return a.sure1x1nice1x1some(a.Done())
}

func (a *Vae[V]) Some() []V {
	return a.sure1x1nice1x1some(a.Done())
}

func (a *Vae[V]) Good() {
	a.good1x1fine1x1safe1x1have(a.Done())
}

func (a *Vae[V]) Fine() {
	a.good1x1fine1x1safe1x1have(a.Done())
}

func (a *Vae[V]) Safe() {
	a.good1x1fine1x1safe1x1have(a.Done())
}

func (a *Vae[V]) Have() {
	a.good1x1fine1x1safe1x1have(a.Done())
}

func (a *Vae[V]) Zero() {
	a.none1x1zero1x1void(a.Done())
}

func (a *Vae[V]) None() {
	a.none1x1zero1x1void(a.Done())
}

func (a *Vae[V]) Void() {
	a.none1x1zero1x1void(a.Done())
}

func (a *Vae[V]) Size(n int) {
	a.size1x1len1x1length(a.Done(), n)
}

func (a *Vae[V]) Len(n int) {
	a.size1x1len1x1length(a.Done(), n)
}

func (a *Vae[V]) Length(n int) {
	a.size1x1len1x1length(a.Done(), n)
}

func (a *Vae[V]) sure1x1nice1x1some(v []V) []V {
	if len(v) == 0 {
		panic(errors.New("SHOULD HAVE SOME BUT IS NONE"))
	}
	return v
}

func (a *Vae[V]) good1x1fine1x1safe1x1have(v []V) {
	if len(v) == 0 {
		panic(errors.New("SHOULD HAVE SOME BUT IS NONE"))
	}
}

func (a *Vae[V]) none1x1zero1x1void(v []V) {
	if len(v) != 0 {
		panic(errors.New("SHOULD BE NONE BUT HAVE SOME"))
	}
}

func (a *Vae[V]) size1x1len1x1length(v []V, n int) {
	if len(v) != n {
		panic(errors.New("SIZE IS NOT SAME"))
	}
}
