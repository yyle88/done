package done

import "github.com/pkg/errors"

// Vae wraps a slice value and an error with slice validation methods.
// Embeds Ve to provide basic error handling methods. A means slice type.
// Vae 封装一个切片值和一个错误，配合切片验证方法。
// 嵌入 Ve 来提供基础错误处理方法。A 代表切片类型。
type Vae[V any] struct {
	V []V
	E error
	*Ve[[]V]
}

// VAE creates a Vae instance with a slice value and error.
// Accepts two params: one is slice and one is error interface.
// VAE 创建一个 Vae 实例，包含切片值和错误。
// 接受两个参数：切片和错误接口。
func VAE[V any](val []V, err error) *Vae[V] {
	return &Vae[V]{
		V:  val,
		E:  err,
		Ve: VE[[]V](val, err),
	}
}

// Sure ensures no error and slice is not empty, then returns the slice.
// Sure 确保无错误且切片非空，然后返回切片。
func (a *Vae[V]) Sure() []V {
	return a.sure1x1nice1x1some(a.Done())
}

// Nice ensures no error and slice is not empty, then returns the slice.
// Nice 确保无错误且切片非空，然后返回切片。
func (a *Vae[V]) Nice() []V {
	return a.sure1x1nice1x1some(a.Done())
}

// Some ensures no error and slice is not empty, then returns the slice.
// Some 确保无错误且切片非空，然后返回切片。
func (a *Vae[V]) Some() []V {
	return a.sure1x1nice1x1some(a.Done())
}

// Good ensures no error and slice is not empty.
// Good 确保无错误且切片非空。
func (a *Vae[V]) Good() {
	a.good1x1fine1x1safe1x1have(a.Done())
}

// Fine ensures no error and slice is not empty.
// Fine 确保无错误且切片非空。
func (a *Vae[V]) Fine() {
	a.good1x1fine1x1safe1x1have(a.Done())
}

// Safe ensures no error and slice is not empty.
// Safe 确保无错误且切片非空。
func (a *Vae[V]) Safe() {
	a.good1x1fine1x1safe1x1have(a.Done())
}

// Have ensures no error and slice is not empty.
// Have 确保无错误且切片非空。
func (a *Vae[V]) Have() {
	a.good1x1fine1x1safe1x1have(a.Done())
}

// Zero ensures no error and slice is empty.
// Zero 确保无错误且切片是空的。
func (a *Vae[V]) Zero() {
	a.none1x1zero1x1void(a.Done())
}

// None ensures no error and slice is empty.
// None 确保无错误且切片是空的。
func (a *Vae[V]) None() {
	a.none1x1zero1x1void(a.Done())
}

// Void ensures no error and slice is empty.
// Void 确保无错误且切片是空的。
func (a *Vae[V]) Void() {
	a.none1x1zero1x1void(a.Done())
}

// Size ensures no error and slice has the given size.
// Size 确保无错误且切片大小等于给定值。
func (a *Vae[V]) Size(n int) {
	a.size1x1len1x1length(a.Done(), n)
}

// Len ensures no error and slice has the given length.
// Len 确保无错误且切片长度等于给定值。
func (a *Vae[V]) Len(n int) {
	a.size1x1len1x1length(a.Done(), n)
}

// Length ensures no error and slice has the given length.
// Length 确保无错误且切片长度等于给定值。
func (a *Vae[V]) Length(n int) {
	a.size1x1len1x1length(a.Done(), n)
}

// sure1x1nice1x1some checks slice is not empty and returns it.
// sure1x1nice1x1some 检查切片非空并返回切片。
func (a *Vae[V]) sure1x1nice1x1some(v []V) []V {
	if len(v) == 0 {
		panic(errors.New("SHOULD HAVE SOME BUT IS NONE"))
	}
	return v
}

// good1x1fine1x1safe1x1have checks slice is not empty.
// good1x1fine1x1safe1x1have 检查切片非空。
func (a *Vae[V]) good1x1fine1x1safe1x1have(v []V) {
	if len(v) == 0 {
		panic(errors.New("SHOULD HAVE SOME BUT IS NONE"))
	}
}

// none1x1zero1x1void checks slice is empty.
// none1x1zero1x1void 检查切片是空的。
func (a *Vae[V]) none1x1zero1x1void(v []V) {
	if len(v) != 0 {
		panic(errors.New("SHOULD BE NONE BUT HAVE SOME"))
	}
}

// size1x1len1x1length checks slice has the given size.
// size1x1len1x1length 检查切片大小等于给定值。
func (a *Vae[V]) size1x1len1x1length(v []V, n int) {
	if len(v) != n {
		panic(errors.New("SIZE IS NOT SAME"))
	}
}
