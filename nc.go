package done

func C0(err error) {
	Must(err)
}

func C1[T1 comparable](v1 T1, err error) T1 {
	Must(err)
	Nice(v1)
	return v1
}

func C2[T1, T2 comparable](v1 T1, v2 T2, err error) (T1, T2) {
	Must(err)
	Nice(v1)
	Nice(v2)
	return v1, v2
}

func C3[T1, T2, T3 comparable](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	Must(err)
	Nice(v1)
	Nice(v2)
	Nice(v3)
	return v1, v2, v3
}

func C4[T1, T2, T3, T4 comparable](v1 T1, v2 T2, v3 T3, v4 T4, err error) (T1, T2, T3, T4) {
	Must(err)
	Nice(v1)
	Nice(v2)
	Nice(v3)
	Nice(v4)
	return v1, v2, v3, v4
}

func C5[T1, T2, T3, T4, T5 comparable](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, err error) (T1, T2, T3, T4, T5) {
	Must(err)
	Nice(v1)
	Nice(v2)
	Nice(v3)
	Nice(v4)
	Nice(v5)
	return v1, v2, v3, v4, v5
}

func C6[T1, T2, T3, T4, T5, T6 comparable](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, err error) (T1, T2, T3, T4, T5, T6) {
	Must(err)
	Nice(v1)
	Nice(v2)
	Nice(v3)
	Nice(v4)
	Nice(v5)
	Nice(v6)
	return v1, v2, v3, v4, v5, v6
}

func C7[T1, T2, T3, T4, T5, T6, T7 comparable](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, err error) (T1, T2, T3, T4, T5, T6, T7) {
	Must(err)
	Nice(v1)
	Nice(v2)
	Nice(v3)
	Nice(v4)
	Nice(v5)
	Nice(v6)
	Nice(v7)
	return v1, v2, v3, v4, v5, v6, v7
}

func C8[T1, T2, T3, T4, T5, T6, T7, T8 comparable](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, err error) (T1, T2, T3, T4, T5, T6, T7, T8) {
	Must(err)
	Nice(v1)
	Nice(v2)
	Nice(v3)
	Nice(v4)
	Nice(v5)
	Nice(v6)
	Nice(v7)
	Nice(v8)
	return v1, v2, v3, v4, v5, v6, v7, v8
}

func C9[T1, T2, T3, T4, T5, T6, T7, T8, T9 comparable](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, err error) (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	Must(err)
	Nice(v1)
	Nice(v2)
	Nice(v3)
	Nice(v4)
	Nice(v5)
	Nice(v6)
	Nice(v7)
	Nice(v8)
	Nice(v9)
	return v1, v2, v3, v4, v5, v6, v7, v8, v9
}
