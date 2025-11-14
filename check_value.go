package done

import "github.com/pkg/errors"

// Sure ensures the value is not zero and returns it.
// Panics if the value equals its zero value.
// Sure 确保值非零并返回该值。
// 如果值等于零值则触发 panic。
func Sure[V comparable](v V) V {
	var zero V
	if v == zero {
		panic(errors.New("SHOULD BE SURE BUT IS ZERO"))
	}
	return v
}

// Nice ensures the value is not zero and returns it.
// Panics if the value equals its zero value.
// Nice 确保值非零并返回该值。
// 如果值等于零值则触发 panic。
func Nice[V comparable](v V) V {
	var zero V
	if v == zero {
		panic(errors.New("SHOULD BE NICE BUT IS ZERO"))
	}
	return v
}

// Good ensures the value is not zero.
// Panics if the value equals its zero value.
// Good 确保值非零。
// 如果值等于零值则触发 panic。
func Good[V comparable](v V) {
	var zero V
	if v == zero {
		panic(errors.New("SHOULD BE GOOD BUT IS ZERO"))
	}
}

// Fine ensures the value is not zero.
// Panics if the value equals its zero value.
// Fine 确保值非零。
// 如果值等于零值则触发 panic。
func Fine[V comparable](v V) {
	var zero V
	if v == zero {
		panic(errors.New("SHOULD BE Fine BUT IS ZERO"))
	}
}

// Safe ensures the value is not zero.
// Panics if the value equals its zero value.
// Safe 确保值非零。
// 如果值等于零值则触发 panic。
func Safe[V comparable](v V) {
	var zero V
	if v == zero {
		panic(errors.New("SHOULD BE SAFE BUT IS ZERO"))
	}
}

// Zero ensures the value is its type's zero value.
// Panics if the value is not zero (e.g., num is not 0, string is not "", pointer is not nil).
// Zero 确保值是其类型的零值。
// 如果值非零则触发 panic（例如数字非0、字符串非空、指针非nil）。
func Zero[V comparable](v V) {
	var zero V
	if v != zero {
		panic(errors.New("SHOULD BE ZERO BUT NOT ZERO"))
	}
}

// None ensures the value is its type's zero value.
// Panics if the value is not zero.
// None 确保值是其类型的零值。
// 如果值非零则触发 panic。
func None[V comparable](v V) {
	var zero V
	if v != zero {
		panic(errors.New("SHOULD BE NONE BUT NOT NONE"))
	}
}

// Null ensures the pointer is nil.
// Panics if the pointer is not nil.
// Null 确保指针是 nil。
// 如果指针非 nil 则触发 panic。
func Null[T any](v *T) {
	if v != nil {
		panic(errors.New("SHOULD BE NULL BUT IS FULL"))
	}
}

// Full ensures the pointer is not nil and returns it.
// Panics if the pointer is nil.
// Full 确保指针非 nil 并返回该指针。
// 如果指针是 nil 则触发 panic。
func Full[T any](v *T) *T {
	if v == nil {
		panic(errors.New("SHOULD BE FULL BUT IS NULL"))
	}
	return v
}
