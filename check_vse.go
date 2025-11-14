package done

import (
	"strings"

	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Vse wraps a string value and an error with string validation methods.
// Embeds Vce to provide comparable validation methods.
// Vse 封装一个字符串值和一个错误，配合字符串验证方法。
// 嵌入 Vce 来提供可比较验证方法。
type Vse struct {
	V string
	E error
	*Vce[string]
}

// VSE creates a Vse instance with a string value and error.
// VSE 创建一个 Vse 实例，包含字符串值和错误。
func VSE(val string, err error) *Vse {
	return &Vse{
		V:   val,
		E:   err,
		Vce: VCE[string](val, err),
	}
}

// Equals ensures no error and string equals the given string.
// Equals 确保无错误且字符串等于给定字符串。
func (a *Vse) Equals(s string) {
	if v := a.Done(); v != s {
		zaplog.LOGS.Skip1.Panic("wrong", zap.String("v", v), zap.String("s", s))
	}
}

// HasPrefix ensures no error and string has the given prefix.
// HasPrefix 确保无错误且字符串包含给定前缀。
func (a *Vse) HasPrefix(prefix string) {
	if v := a.Done(); !strings.HasPrefix(v, prefix) {
		zaplog.LOGS.Skip1.Panic("wrong", zap.String("v", v), zap.String("prefix", prefix))
	}
}

// HasSuffix ensures no error and string has the given suffix.
// HasSuffix 确保无错误且字符串包含给定后缀。
func (a *Vse) HasSuffix(suffix string) {
	if v := a.Done(); !strings.HasSuffix(v, suffix) {
		zaplog.LOGS.Skip1.Panic("wrong", zap.String("v", v), zap.String("suffix", suffix))
	}
}

// Contains ensures no error and string contains the given substring.
// Contains 确保无错误且字符串包含给定子串。
func (a *Vse) Contains(sub string) {
	if v := a.Done(); !strings.Contains(v, sub) {
		zaplog.LOGS.Skip1.Panic("wrong", zap.String("v", v), zap.String("sub", sub))
	}
}
