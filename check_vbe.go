package done

import "github.com/pkg/errors"

// Vbe wraps a boolean value and an error with boolean validation methods.
// Embeds Vce to provide comparable validation methods.
// Vbe 封装一个布尔值和一个错误，配合布尔验证方法。
// 嵌入 Vce 来提供可比较验证方法。
type Vbe struct {
	V bool
	E error
	*Vce[bool]
}

// VBE creates a Vbe instance with a boolean value and error.
// VBE 创建一个 Vbe 实例，包含布尔值和错误。
func VBE(val bool, err error) *Vbe {
	return &Vbe{
		V:   val,
		E:   err,
		Vce: VCE[bool](val, err),
	}
}

// TRUE ensures no error and value is true.
// TRUE 确保无错误且值是 true。
func (a *Vbe) TRUE() {
	a.assert1T1(a.Done())
}

// FALSE ensures no error and value is false.
// FALSE 确保无错误且值是 false。
func (a *Vbe) FALSE() {
	a.assert0F0(a.Done())
}

// Yes ensures no error and value is true.
// Yes 确保无错误且值是 true。
func (a *Vbe) Yes() {
	a.assert1T1(a.Done())
}

// YES ensures no error and value is true.
// YES 确保无错误且值是 true。
func (a *Vbe) YES() {
	a.assert1T1(a.Done())
}

// NO ensures no error and value is false.
// NO 确保无错误且值是 false。
func (a *Vbe) NO() {
	a.assert0F0(a.Done())
}

// No ensures no error and value is false.
// No 确保无错误且值是 false。
func (a *Vbe) No() {
	a.assert0F0(a.Done())
}

// OK ensures no error and value is true.
// OK 确保无错误且值是 true。
func (a *Vbe) OK() {
	a.assert1T1(a.Done())
}

// Not ensures no error and value is false.
// Not 确保无错误且值是 false。
func (a *Vbe) Not() {
	a.assert0F0(a.Done())
}

// assert1T1 panics if value is not true.
// assert1T1 在值不是 true 时触发 panic。
func (a *Vbe) assert1T1(v bool) {
	if !v {
		panic(errors.New("SHOULD BE TRUE BUT IS FALSE"))
	}
}

// assert0F0 panics if value is not false.
// assert0F0 在值不是 false 时触发 panic。
func (a *Vbe) assert0F0(v bool) {
	if v {
		panic(errors.New("SHOULD BE FALSE BUT IS TRUE"))
	}
}
