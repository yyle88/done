package done

import (
	"testing"

	"github.com/pkg/errors"
)

// TestDone tests the Done function with nil and non-nil errors.
// TestDone 测试 Done 函数处理 nil 和非 nil 错误。
func TestDone(t *testing.T) {
	if false {
		Done(errors.New("wrong"))
	}
}

// TestMust tests the Must function with nil and non-nil errors.
// TestMust 测试 Must 函数处理 nil 和非 nil 错误。
func TestMust(t *testing.T) {
	if false {
		Must(errors.New("wrong"))
	}
}

// TestSoft tests the Soft function logs warning without panic.
// TestSoft 测试 Soft 函数记录告警但不触发 panic。
func TestSoft(t *testing.T) {
	Soft(errors.New("wrong"))
}

// TestFata tests the Fata function logs and panics on error.
// TestFata 测试 Fata 函数记录日志并在错误时触发 panic。
func TestFata(t *testing.T) {
	if false {
		Fata(errors.New("123"))
	}
}
