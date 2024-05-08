package done

type Ve[V any] struct {
	V V
	E error
}

func VE[V any](val V, err error) *Ve[V] {
	return &Ve[V]{V: val, E: err}
}

func (a *Ve[V]) Done() V {
	Done(a.E)
	return a.V
}
