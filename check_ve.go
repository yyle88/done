package done

type Ve[V any] struct {
	V V
	E error
}

// VE accept two params. one is any and one is error interface
func VE[V any](val V, err error) *Ve[V] {
	return &Ve[V]{
		V: val,
		E: err,
	}
}

func (a *Ve[V]) Value() V {
	Done(a.E)  //只会检查错误，而不检查结果
	return a.V //结果不做检查，允许返回零值，因为不是 comparable 的这里也没法比较是否为零值
}

func (a *Ve[V]) Done() V {
	Done(a.E)  //只会检查错误，而不检查结果
	return a.V //结果不做检查，允许返回零值，因为不是 comparable 的这里也没法比较是否为零值
}

func (a *Ve[V]) Must() V {
	Done(a.E)  //只会检查错误，而不检查结果
	return a.V //结果不做检查，允许返回零值
}

func (a *Ve[V]) Soft() V {
	Soft(a.E)  //有错误时打印告警日志
	return a.V //结果不做检查直接返回
}

func (a *Ve[V]) Omit() V {
	return a.V //即使出错也不报错，只把结果返回，结果有可能是正常值也有可能是零值或异常值
}

func (a *Ve[V]) Skip() V {
	return a.V //即使出错也不报错，只把结果返回，结果有可能是正常值也有可能是零值或异常值
}
