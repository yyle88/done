package done

func V0(err error) {
	Must(err)
}

func V1[T1 any](v1 T1, err error) T1 {
	Must(err)
	return v1
}

func V2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	Must(err)
	return v1, v2
}

func V3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	Must(err)
	return v1, v2, v3
}

func V4[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4, err error) (T1, T2, T3, T4) {
	Must(err)
	return v1, v2, v3, v4
}

func V5[T1, T2, T3, T4, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, err error) (T1, T2, T3, T4, T5) {
	Must(err)
	return v1, v2, v3, v4, v5
}

func V6[T1, T2, T3, T4, T5, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, err error) (T1, T2, T3, T4, T5, T6) {
	Must(err)
	return v1, v2, v3, v4, v5, v6
}

func V7[T1, T2, T3, T4, T5, T6, T7 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, err error) (T1, T2, T3, T4, T5, T6, T7) {
	Must(err)
	return v1, v2, v3, v4, v5, v6, v7
}

func V8[T1, T2, T3, T4, T5, T6, T7, T8 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, err error) (T1, T2, T3, T4, T5, T6, T7, T8) {
	Must(err)
	return v1, v2, v3, v4, v5, v6, v7, v8
}

func V9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, err error) (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	Must(err)
	return v1, v2, v3, v4, v5, v6, v7, v8, v9
}
