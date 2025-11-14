package done

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestGood tests the Good function with non-zero values.
// TestGood 测试 Good 函数处理非零值。
func TestGood(t *testing.T) {
	Good(&Example{})
	Good(Example{S: "xyz"})
}

// TestNice tests the Nice function returns non-zero values.
// TestNice 测试 Nice 函数返回非零值。
func TestNice(t *testing.T) {
	a := Nice(Example{S: "abc"})
	require.Equal(t, "abc", a.S)
	p := Nice(&Example{S: "uvw"})
	require.Equal(t, "uvw", p.S)
}

// TestZero tests the Zero function with zero values.
// TestZero 测试 Zero 函数处理零值。
func TestZero(t *testing.T) {
	Zero((*Example)(nil))
	Zero(Example{})
}

// Example is a test struct with a string field.
// Example 是包含字符串字段的测试结构体。
type Example struct {
	S string
}

// newExample creates a new Example instance.
// newExample 创建一个新的 Example 实例。
func newExample() (*Example, error) {
	return &Example{S: "xyz"}, nil
}

// TestNull tests the Null function ensures pointer is nil.
// TestNull 测试 Null 函数确保指针是 nil。
func TestNull(t *testing.T) {
	var example *Example
	Null(example)

	require.Panics(t, func() {
		Null(&Example{S: "abc"})
	})
}

// TestFull tests the Full function ensures pointer is not nil.
// TestFull 测试 Full 函数确保指针非 nil。
func TestFull(t *testing.T) {
	Full(&Example{S: "abc"})

	require.Panics(t, func() {
		var example *Example
		Full(example)
	})
}
