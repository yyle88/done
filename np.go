package done

func P0(err error) {
	Must(err)
}

func P1[T1 any](v1 *T1, err error) *T1 {
	Must(err)
	Full(v1)
	return v1
}

func P2[T1, T2 any](v1 *T1, v2 *T2, err error) (*T1, *T2) {
	Must(err)
	Full(v1)
	Full(v2)
	return v1, v2
}

func P3[T1, T2, T3 any](v1 *T1, v2 *T2, v3 *T3, err error) (*T1, *T2, *T3) {
	Must(err)
	Full(v1)
	Full(v2)
	Full(v3)
	return v1, v2, v3
}

func P4[T1, T2, T3, T4 any](v1 *T1, v2 *T2, v3 *T3, v4 *T4, err error) (*T1, *T2, *T3, *T4) {
	Must(err)
	Full(v1)
	Full(v2)
	Full(v3)
	Full(v4)
	return v1, v2, v3, v4
}

func P5[T1, T2, T3, T4, T5 any](v1 *T1, v2 *T2, v3 *T3, v4 *T4, v5 *T5, err error) (*T1, *T2, *T3, *T4, *T5) {
	Must(err)
	Full(v1)
	Full(v2)
	Full(v3)
	Full(v4)
	Full(v5)
	return v1, v2, v3, v4, v5
}

func P6[T1, T2, T3, T4, T5, T6 any](v1 *T1, v2 *T2, v3 *T3, v4 *T4, v5 *T5, v6 *T6, err error) (*T1, *T2, *T3, *T4, *T5, *T6) {
	Must(err)
	Full(v1)
	Full(v2)
	Full(v3)
	Full(v4)
	Full(v5)
	Full(v6)
	return v1, v2, v3, v4, v5, v6
}

func P7[T1, T2, T3, T4, T5, T6, T7 any](v1 *T1, v2 *T2, v3 *T3, v4 *T4, v5 *T5, v6 *T6, v7 *T7, err error) (*T1, *T2, *T3, *T4, *T5, *T6, *T7) {
	Must(err)
	Full(v1)
	Full(v2)
	Full(v3)
	Full(v4)
	Full(v5)
	Full(v6)
	Full(v7)
	return v1, v2, v3, v4, v5, v6, v7
}

func P8[T1, T2, T3, T4, T5, T6, T7, T8 any](v1 *T1, v2 *T2, v3 *T3, v4 *T4, v5 *T5, v6 *T6, v7 *T7, v8 *T8, err error) (*T1, *T2, *T3, *T4, *T5, *T6, *T7, *T8) {
	Must(err)
	Full(v1)
	Full(v2)
	Full(v3)
	Full(v4)
	Full(v5)
	Full(v6)
	Full(v7)
	Full(v8)
	return v1, v2, v3, v4, v5, v6, v7, v8
}

func P9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](v1 *T1, v2 *T2, v3 *T3, v4 *T4, v5 *T5, v6 *T6, v7 *T7, v8 *T8, v9 *T9, err error) (*T1, *T2, *T3, *T4, *T5, *T6, *T7, *T8, *T9) {
	Must(err)
	Full(v1)
	Full(v2)
	Full(v3)
	Full(v4)
	Full(v5)
	Full(v6)
	Full(v7)
	Full(v8)
	Full(v9)
	return v1, v2, v3, v4, v5, v6, v7, v8, v9
}
