package done

import "github.com/pkg/errors"

// Vme wraps a map value and an error with map validation methods.
// Embeds Ve to provide basic error handling methods. M means map type.
// Vme 封装一个 map 值和一个错误，配合 map 验证方法。
// 嵌入 Ve 来提供基础错误处理方法。M 代表 map 类型。
type Vme[K comparable, V any] struct {
	V map[K]V
	E error
	*Ve[map[K]V]
}

// VME creates a Vme instance with a map value and error.
// Accepts two params: one is map and one is error interface.
// VME 创建一个 Vme 实例，包含 map 值和错误。
// 接受两个参数：map 和错误接口。
func VME[K comparable, V any](val map[K]V, err error) *Vme[K, V] {
	return &Vme[K, V]{
		V:  val,
		E:  err,
		Ve: VE[map[K]V](val, err),
	}
}

// Sure ensures no error and map is not empty, then returns the map.
// Sure 确保无错误且 map 非空，然后返回 map。
func (a *Vme[K, V]) Sure() map[K]V {
	return a.sure1x1nice1x1some(a.Done())
}

// Nice ensures no error and map is not empty, then returns the map.
// Nice 确保无错误且 map 非空，然后返回 map。
func (a *Vme[K, V]) Nice() map[K]V {
	return a.sure1x1nice1x1some(a.Done())
}

// Some ensures no error and map is not empty, then returns the map.
// Some 确保无错误且 map 非空，然后返回 map。
func (a *Vme[K, V]) Some() map[K]V {
	return a.sure1x1nice1x1some(a.Done())
}

// Good ensures no error and map is not empty.
// Good 确保无错误且 map 非空。
func (a *Vme[K, V]) Good() {
	a.good1x1fine1x1safe1x1have(a.Done())
}

// Fine ensures no error and map is not empty.
// Fine 确保无错误且 map 非空。
func (a *Vme[K, V]) Fine() {
	a.good1x1fine1x1safe1x1have(a.Done())
}

// Safe ensures no error and map is not empty.
// Safe 确保无错误且 map 非空。
func (a *Vme[K, V]) Safe() {
	a.good1x1fine1x1safe1x1have(a.Done())
}

// Have ensures no error and map is not empty.
// Have 确保无错误且 map 非空。
func (a *Vme[K, V]) Have() {
	a.good1x1fine1x1safe1x1have(a.Done())
}

// Zero ensures no error and map is empty.
// Zero 确保无错误且 map 是空的。
func (a *Vme[K, V]) Zero() {
	a.none1x1zero1x1void(a.Done())
}

// None ensures no error and map is empty.
// None 确保无错误且 map 是空的。
func (a *Vme[K, V]) None() {
	a.none1x1zero1x1void(a.Done())
}

// Void ensures no error and map is empty.
// Void 确保无错误且 map 是空的。
func (a *Vme[K, V]) Void() {
	a.none1x1zero1x1void(a.Done())
}

// Size ensures no error and map has the given size.
// Size 确保无错误且 map 大小等于给定值。
func (a *Vme[K, V]) Size(n int) {
	a.size1x1len1x1length(a.Done(), n)
}

// Len ensures no error and map has the given length.
// Len 确保无错误且 map 长度等于给定值。
func (a *Vme[K, V]) Len(n int) {
	a.size1x1len1x1length(a.Done(), n)
}

// Length ensures no error and map has the given length.
// Length 确保无错误且 map 长度等于给定值。
func (a *Vme[K, V]) Length(n int) {
	a.size1x1len1x1length(a.Done(), n)
}

// sure1x1nice1x1some checks map is not empty and returns it.
// sure1x1nice1x1some 检查 map 非空并返回 map。
func (a *Vme[K, V]) sure1x1nice1x1some(v map[K]V) map[K]V {
	if len(v) == 0 {
		panic(errors.New("SHOULD HAVE SOME BUT IS NONE"))
	}
	return v
}

// good1x1fine1x1safe1x1have checks map is not empty.
// good1x1fine1x1safe1x1have 检查 map 非空。
func (a *Vme[K, V]) good1x1fine1x1safe1x1have(v map[K]V) {
	if len(v) == 0 {
		panic(errors.New("SHOULD HAVE SOME BUT IS NONE"))
	}
}

// none1x1zero1x1void checks map is empty.
// none1x1zero1x1void 检查 map 是空的。
func (a *Vme[K, V]) none1x1zero1x1void(v map[K]V) {
	if len(v) != 0 {
		panic(errors.New("SHOULD BE NONE BUT HAVE SOME"))
	}
}

// size1x1len1x1length checks map has the given size.
// size1x1len1x1length 检查 map 大小等于给定值。
func (a *Vme[K, V]) size1x1len1x1length(v map[K]V, n int) {
	if len(v) != n {
		panic(errors.New("SIZE IS NOT SAME"))
	}
}
