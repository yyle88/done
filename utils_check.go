package done

import (
	"fmt"
	"os"

	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func Done(err error) {
	if err != nil {
		zaplog.LOGS.P1.Panic("ERROR", zap.Error(err))
	}
}

func Must(err error) {
	if err != nil {
		zaplog.LOGS.P1.Panic("ERROR", zap.Error(err))
	}
}

func Soft(err error) {
	if err != nil {
		zaplog.LOGS.P1.Warn("WARN", zap.Error(err))
	}
}

// Fata FaFaFaFaFa 哈哈哈哈哈
func Fata(err error) {
	if err != nil {
		zaplog.LOGS.P1.Fatal("ERROR", zap.Error(err))
	}
}

// EExt if err != nil os.Exit(code)
func EExt(err error, code int) {
	if err != nil {
		var msg = fmt.Sprintf("\x1b[31;1m%s\x1b[0m", err) //红色打印
		fmt.Println(msg)
		zaplog.LOGS.P1.Error("ERROR", zap.Error(err))
		os.Exit(code)
	}
}
