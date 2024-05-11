package done

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func Done(err error) {
	if err != nil {
		zaplog.LOGS.P1.Panic("ERROR", zap.Error(errors.WithMessage(err, "wrong")))
	}
}

func Must(err error, msg string) {
	if err != nil {
		zaplog.LOGS.P1.Panic("ERROR", zap.Error(errors.WithMessage(err, msg)))
	}
}

func Soft(err error, msg string) {
	if err != nil {
		zaplog.LOGS.P1.Warn("ERROR", zap.Error(errors.WithMessage(err, msg)))
	}
}

// Fata FaFaFaFaFa 哈哈哈哈哈
func Fata(err error, msg string) {
	if err != nil {
		zaplog.LOGS.P1.Fatal("ERROR", zap.Error(errors.WithMessage(err, msg)))
	}
}

func EExt(err error, code int) {
	if err != nil {
		const format = "\x1b[31;1m%s\x1b[0m" //红色打印
		var message = fmt.Sprintf(format, errors.WithMessage(err, "wrong"))
		fmt.Println(message)
		zaplog.LOGS.P1.Warn("ERROR", zap.Error(errors.WithMessage(err, "wrong")))
		os.Exit(code)
	}
}
