package done

import "github.com/pkg/errors"

// numType defines constraint for numeric types.
// numType 定义数字类型约束。
type numType interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | uint | float32 | float64
}

// Vne wraps a numeric value and an error with numeric validation methods.
// Embeds Vce to provide comparable validation methods.
// Vne 封装一个数字值和一个错误，配合数字验证方法。
// 嵌入 Vce 来提供可比较验证方法。
type Vne[V numType] struct {
	V V
	E error
	*Vce[V]
}

// VNE creates a Vne instance with a numeric value and error.
// VNE 创建一个 Vne 实例，包含数字值和错误。
func VNE[V numType](val V, err error) *Vne[V] {
	return &Vne[V]{
		V:   val,
		E:   err,
		Vce: VCE[V](val, err),
	}
}

// Gt ensures no error and value is greater than base, then returns the value.
// Gt 确保无错误且值大于基准值，然后返回值。
func (a *Vne[V]) Gt(base V) V {
	if x := a.Done(); x > base {
		return x
	} else {
		panic(errors.Errorf("SHOULD BE x > BASE BUT NOT. x=%v BASE=%v", x, base))
	}
}

// Lt ensures no error and value is less than base, then returns the value.
// Lt 确保无错误且值小于基准值，然后返回值。
func (a *Vne[V]) Lt(base V) V {
	if x := a.Done(); x < base {
		return x
	} else {
		panic(errors.Errorf("SHOULD BE x < BASE BUT NOT. x=%v BASE=%v", x, base))
	}
}

// Gte ensures no error and value is greater than or equal to base, then returns the value.
// Gte 确保无错误且值大于等于基准值，然后返回值。
func (a *Vne[V]) Gte(base V) V {
	if x := a.Done(); x >= base {
		return x
	} else {
		panic(errors.Errorf("SHOULD BE x >= BASE BUT NOT. x=%v BASE=%v", x, base))
	}
}

// Lte ensures no error and value is less than or equal to base, then returns the value.
// Lte 确保无错误且值小于等于基准值，然后返回值。
func (a *Vne[V]) Lte(base V) V {
	if x := a.Done(); x <= base {
		return x
	} else {
		panic(errors.Errorf("SHOULD BE x <= BASE BUT NOT. x=%v BASE=%v", x, base))
	}
}
