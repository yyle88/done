package done

// Ve wraps a value and an error with validation methods.
// Ve 封装一个值和一个错误，配合验证方法。
type Ve[V any] struct {
	V V
	E error
}

// VE creates a Ve instance with a value and error.
// Accepts two params: one is any value and one is error interface.
// VE 创建一个 Ve 实例，包含值和错误。
// 接受两个参数：任意值和错误接口。
func VE[V any](val V, err error) *Ve[V] {
	return &Ve[V]{
		V: val,
		E: err,
	}
}

// Done checks if error exists and panics, then returns the value.
// The value is not checked and can be zero.
// Done 检查是否存在错误并触发 panic，然后返回值。
// 不检查值，可以是零值。
func (a *Ve[V]) Done() V {
	Done(a.E)  // Check error and panic if exists // 检查错误，存在则 panic
	return a.V // Return value without checking // 返回值但不检查
}

// Must checks if error exists and panics, then returns the value.
// The value is not checked and can be zero.
// Must 检查是否存在错误并触发 panic，然后返回值。
// 不检查值，可以是零值。
func (a *Ve[V]) Must() V {
	Done(a.E)  // Check error and panic if exists // 检查错误，存在则 panic
	return a.V // Return value without checking // 返回值但不检查
}

// Soft logs a warning if error exists, then returns the value.
// Returns the value without checking.
// Soft 如果存在错误则记录告警，然后返回值。
// 返回值但不检查。
func (a *Ve[V]) Soft() V {
	Soft(a.E)  // Log warning if error exists // 有错误时记录告警
	return a.V // Return value without checking // 返回值但不检查
}

// Omit returns the value and ignores the error.
// The value can be normal, zero, or unexpected.
// Omit 返回值并忽视错误。
// 值可能是正常值、零值或异常值。
func (a *Ve[V]) Omit() V {
	return a.V // Return value even if error exists // 即使有错误也返回值
}

// Skip returns the value and ignores the error.
// The value can be normal, zero, or unexpected.
// Skip 返回值并忽视错误。
// 值可能是正常值、零值或异常值。
func (a *Ve[V]) Skip() V {
	return a.V // Return value even if error exists // 即使有错误也返回值
}
