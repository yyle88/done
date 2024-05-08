package done

import "github.com/pkg/errors"

func Sure[V comparable](v V) V {
	var zero V
	if !(v != zero) {
		panic(errors.New("SHOULD BE SURE BUT IS ZERO"))
	}
	return v
}

func Nice[V comparable](v V) V {
	var zero V
	if !(v != zero) {
		panic(errors.New("SHOULD BE NICE BUT IS ZERO"))
	}
	return v
}

func Good[V comparable](v V) {
	var zero V
	if !(v != zero) {
		panic(errors.New("SHOULD BE GOOD BUT IS ZERO"))
	}
}

func Fine[V comparable](v V) {
	var zero V
	if !(v != zero) {
		panic(errors.New("SHOULD BE Fine BUT IS ZERO"))
	}
}

func Safe[V comparable](v V) {
	var zero V
	if !(v != zero) {
		panic(errors.New("SHOULD BE SAFE BUT IS ZERO"))
	}
}

func Zero[V comparable](v V) {
	var zero V
	if !(v == zero) {
		panic(errors.New("SHOULD BE ZERO BUT NOT ZERO"))
	}
}

func None[V comparable](v V) {
	var zero V
	if !(v == zero) {
		panic(errors.New("SHOULD BE NONE BUT NOT NONE"))
	}
}
