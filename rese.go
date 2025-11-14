package done

// V0 checks error and panics if exists.
// V0 检查错误，存在则触发 panic。
func V0(err error) {
	Done(err)
}

// V1 checks error and returns one value.
// V1 检查错误并返回一个值。
func V1[T1 any](v1 T1, err error) T1 {
	Done(err)
	return v1
}

// V2 checks error and returns two values.
// V2 检查错误并返回两个值。
func V2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	Done(err)
	return v1, v2
}

// P0 checks error and panics if exists.
// P0 检查错误，存在则触发 panic。
func P0(err error) {
	Done(err)
}

// P1 checks error, ensures pointer is not nil, and returns the pointer.
// P1 检查错误，确保指针非 nil，并返回指针。
func P1[T1 any](v1 *T1, err error) *T1 {
	Done(err)
	Full(v1)
	return v1
}

// P2 checks error, ensures pointers are not nil, and returns the pointers.
// P2 检查错误，确保指针非 nil，并返回指针。
func P2[T1, T2 any](v1 *T1, v2 *T2, err error) (*T1, *T2) {
	Done(err)
	Full(v1)
	Full(v2)
	return v1, v2
}

// C0 checks error and panics if exists.
// C0 检查错误，存在则触发 panic。
func C0(err error) {
	Done(err)
}

// C1 checks error, ensures value is not zero, and returns the value.
// C1 检查错误，确保值非零，并返回值。
func C1[T1 comparable](v1 T1, err error) T1 {
	Done(err)
	Nice(v1)
	return v1
}

// C2 checks error, ensures values are not zero, and returns the values.
// C2 检查错误，确保值非零，并返回值。
func C2[T1, T2 comparable](v1 T1, v2 T2, err error) (T1, T2) {
	Done(err)
	Nice(v1)
	Nice(v2)
	return v1, v2
}
