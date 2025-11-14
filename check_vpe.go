package done

// Vpe wraps a pointer value and an error with validation methods.
// Embeds Ve to provide basic error handling methods.
// Vpe 封装一个指针值和一个错误，配合验证方法。
// 嵌入 Ve 来提供基础错误处理方法。
type Vpe[V any] struct {
	V       *V
	E       error
	*Ve[*V] // Embed Ve to validate pointers // 嵌入 Ve 来验证指针
}

// VPE creates a Vpe instance with a pointer value and error.
// Accepts two params: one is pointer and one is error interface.
// VPE 创建一个 Vpe 实例，包含指针值和错误。
// 接受两个参数：指针和错误接口。
func VPE[V any](val *V, err error) *Vpe[V] {
	return &Vpe[V]{
		V:  val,
		E:  err,
		Ve: VE[*V](val, err),
	}
}

// Sure ensures no error and pointer is not nil, then returns the pointer.
// Sure 确保无错误且指针非 nil，然后返回指针。
func (a *Vpe[V]) Sure() *V {
	return Sure(a.Done())
}

// Nice ensures no error and pointer is not nil, then returns the pointer.
// Nice 确保无错误且指针非 nil，然后返回指针。
func (a *Vpe[V]) Nice() *V {
	return Nice(a.Done())
}

// Good ensures no error and pointer is not nil.
// Good 确保无错误且指针非 nil。
func (a *Vpe[V]) Good() {
	Good(a.Done())
}

// Fine ensures no error and pointer is not nil.
// Fine 确保无错误且指针非 nil。
func (a *Vpe[V]) Fine() {
	Fine(a.Done())
}

// Safe ensures no error and pointer is not nil.
// Safe 确保无错误且指针非 nil。
func (a *Vpe[V]) Safe() {
	Safe(a.Done())
}

// Zero ensures no error and pointer is nil.
// Zero 确保无错误且指针是 nil。
func (a *Vpe[V]) Zero() {
	Zero(a.Done())
}

// None ensures no error and pointer is nil.
// None 确保无错误且指针是 nil。
func (a *Vpe[V]) None() {
	None(a.Done())
}

// Null ensures no error and pointer is nil.
// Null 确保无错误且指针是 nil。
func (a *Vpe[V]) Null() {
	Null(a.Done())
}

// Full ensures no error and pointer is not nil, then returns the pointer.
// Full 确保无错误且指针非 nil，然后返回指针。
func (a *Vpe[V]) Full() *V {
	return Full(a.Done())
}
