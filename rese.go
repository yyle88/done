package done

func V0(err error) {
	Done(err)
}

func V1[T1 any](v1 T1, err error) T1 {
	Done(err)
	return v1
}

func V2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	Done(err)
	return v1, v2
}

func P0(err error) {
	Done(err)
}

func P1[T1 any](v1 *T1, err error) *T1 {
	Done(err)
	Full(v1)
	return v1
}

func P2[T1, T2 any](v1 *T1, v2 *T2, err error) (*T1, *T2) {
	Done(err)
	Full(v1)
	Full(v2)
	return v1, v2
}

func C0(err error) {
	Done(err)
}

func C1[T1 comparable](v1 T1, err error) T1 {
	Done(err)
	Nice(v1)
	return v1
}

func C2[T1, T2 comparable](v1 T1, v2 T2, err error) (T1, T2) {
	Done(err)
	Nice(v1)
	Nice(v2)
	return v1, v2
}
