package done

import "github.com/pkg/errors"

// Vce wraps a comparable value and an error with validation methods.
// Embeds Ve to provide basic error handling methods.
// Vce 封装一个可比较值和一个错误，配合验证方法。
// 嵌入 Ve 来提供基础错误处理方法。
type Vce[V comparable] struct {
	V V
	E error
	*Ve[V]
}

// VCE creates a Vce instance with a comparable value and error.
// Accepts two params: one is comparable value and one is error interface.
// VCE 创建一个 Vce 实例，包含可比较值和错误。
// 接受两个参数：可比较值和错误接口。
func VCE[V comparable](val V, err error) *Vce[V] {
	return &Vce[V]{
		V:  val,
		E:  err,
		Ve: VE[V](val, err),
	}
}

// Sure ensures no error and value is not zero, then returns the value.
// Sure 确保无错误且值非零，然后返回值。
func (a *Vce[V]) Sure() V {
	return Sure(a.Done())
}

// Nice ensures no error and value is not zero, then returns the value.
// Nice 确保无错误且值非零，然后返回值。
func (a *Vce[V]) Nice() V {
	return Nice(a.Done())
}

// Good ensures no error and value is not zero.
// Good 确保无错误且值非零。
func (a *Vce[V]) Good() {
	Good(a.Done())
}

// Fine ensures no error and value is not zero.
// Fine 确保无错误且值非零。
func (a *Vce[V]) Fine() {
	Fine(a.Done())
}

// Safe ensures no error and value is not zero.
// Safe 确保无错误且值非零。
func (a *Vce[V]) Safe() {
	Safe(a.Done())
}

// Zero ensures no error and value is zero.
// Zero 确保无错误且值是零值。
func (a *Vce[V]) Zero() {
	Zero(a.Done())
}

// None ensures no error and value is zero.
// None 确保无错误且值是零值。
func (a *Vce[V]) None() {
	None(a.Done())
}

// Same ensures no error and value equals the given value.
// Same 确保无错误且值等于给定值。
func (a *Vce[V]) Same(v V) {
	if x := a.Done(); x != v {
		panic(errors.New("SHOULD BE SAME BUT IS DIFFERENT"))
	}
}

// Diff ensures no error and value is not equal to the given value.
// Diff 确保无错误且值不等于给定值。
func (a *Vce[V]) Diff(v V) {
	if x := a.Done(); x == v {
		panic(errors.New("SHOULD BE DIFFERENT BUT IS SAME"))
	}
}

// Is ensures no error and value equals the given value.
// Is 确保无错误且值等于给定值。
func (a *Vce[V]) Is(v V) {
	if x := a.Done(); x != v {
		panic(errors.New("SHOULD BE SAME BUT IS DIFFERENT"))
	}
}

// Equals ensures no error and value equals the given value.
// Equals 确保无错误且值等于给定值。
func (a *Vce[V]) Equals(v V) {
	if x := a.Done(); x != v {
		panic(errors.New("SHOULD BE SAME BUT IS DIFFERENT"))
	}
}

// Different ensures no error and value is not equal to the given value.
// Different 确保无错误且值不等于给定值。
func (a *Vce[V]) Different(v V) {
	if x := a.Done(); x == v {
		panic(errors.New("SHOULD BE DIFFERENT BUT IS SAME"))
	}
}
