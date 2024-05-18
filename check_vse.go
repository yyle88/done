package done

import (
	"strings"

	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type Vse struct {
	V string
	E error
	*Vce[string]
}

func VSE(val string, err error) *Vse {
	return &Vse{
		V:   val,
		E:   err,
		Vce: VCE[string](val, err),
	}
}

func (a *Vse) Equals(s string) {
	if v := a.Done(); v != s {
		zaplog.LOGS.P1.Panic("wrong", zap.String("v", v), zap.String("s", s))
	}
}

func (a *Vse) HasPrefix(prefix string) {
	if v := a.Done(); !strings.HasPrefix(v, prefix) {
		zaplog.LOGS.P1.Panic("wrong", zap.String("v", v), zap.String("prefix", prefix))
	}
}

func (a *Vse) HasSuffix(suffix string) {
	if v := a.Done(); !strings.HasSuffix(v, suffix) {
		zaplog.LOGS.P1.Panic("wrong", zap.String("v", v), zap.String("suffix", suffix))
	}
}

func (a *Vse) Contains(sub string) {
	if v := a.Done(); !strings.Contains(v, sub) {
		zaplog.LOGS.P1.Panic("wrong", zap.String("v", v), zap.String("sub", sub))
	}
}
