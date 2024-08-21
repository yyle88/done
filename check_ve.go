package done

type Ve[V any] struct {
	V V
	E error
}

func VE[V any](val V, err error) *Ve[V] {
	return &Ve[V]{
		V: val,
		E: err,
	}
}

func (a *Ve[V]) Done() V {
	Done(a.E)  //只会检查错误，而不检查结果，因此这个类不要定义Must函数，否则会出问题
	return a.V //结果不做检查
}

// 这个类不要定义Must函数，否则会出问题
//func (a *Ve[V]) Must() V {
//	Done(a.E)  //只会检查错误，而不检查结果，因此这个类不要定义Must函数，否则会出问题
//	return a.V //结果不做检查
//}

func (a *Ve[V]) Soft() V {
	Soft(a.E)  //有错误时打印告警日志
	return a.V //结果不做检查
}

func (a *Ve[V]) Omit() V {
	return a.V //即使出错也不报错，只把结果返回，结果有可能是正常值也有可能是零值或异常值
}

func (a *Ve[V]) Skip() V {
	return a.V //即使出错也不报错，只把结果返回，结果有可能是正常值也有可能是零值或异常值
}
