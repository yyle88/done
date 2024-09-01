package done

import "github.com/pkg/errors"

// Sure means the value is not zero value, then return value
func Sure[V comparable](v V) V {
	var zero V
	if !(v != zero) {
		panic(errors.New("SHOULD BE SURE BUT IS ZERO"))
	}
	return v
}

// Nice means the value is not zero value, then return value
func Nice[V comparable](v V) V {
	var zero V
	if !(v != zero) {
		panic(errors.New("SHOULD BE NICE BUT IS ZERO"))
	}
	return v
}

// Good means the value is not zero value, then return value
func Good[V comparable](v V) {
	var zero V
	if !(v != zero) {
		panic(errors.New("SHOULD BE GOOD BUT IS ZERO"))
	}
}

// Fine means the value is not zero value
func Fine[V comparable](v V) {
	var zero V
	if !(v != zero) {
		panic(errors.New("SHOULD BE Fine BUT IS ZERO"))
	}
}

// Safe means the value is not zero value
func Safe[V comparable](v V) {
	var zero V
	if !(v != zero) {
		panic(errors.New("SHOULD BE SAFE BUT IS ZERO"))
	}
}

// Zero means the value is a zero value, num is 0, string is "", PTR is none.
func Zero[V comparable](v V) {
	var zero V
	if !(v == zero) {
		panic(errors.New("SHOULD BE ZERO BUT NOT ZERO"))
	}
}

// None means the value is a zero value
func None[V comparable](v V) {
	var zero V
	if !(v == zero) {
		panic(errors.New("SHOULD BE NONE BUT NOT NONE"))
	}
}
