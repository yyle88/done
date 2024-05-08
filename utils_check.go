package done

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"gitlab.yyle.com/golang/uvyyle.git/utils_logs/utils_logzap"
	"go.uber.org/zap"
)

func Done(err error) {
	if err != nil {
		utils_logzap.ZAPS.P1.LOG.Panic("ERROR", zap.Error(errors.WithMessage(err, "wrong")))
	}
}

func Must(err error, msg string) {
	if err != nil {
		utils_logzap.ZAPS.P1.LOG.Panic("ERROR", zap.Error(errors.WithMessage(err, msg)))
	}
}

func Soft(err error, msg string) {
	if err != nil {
		utils_logzap.ZAPS.P1.LOG.Warn("ERROR", zap.Error(errors.WithMessage(err, msg)))
	}
}

// Fata FaFaFaFaFa 哈哈哈哈哈
func Fata(err error, msg string) {
	if err != nil {
		utils_logzap.ZAPS.P1.LOG.Fatal("ERROR", zap.Error(errors.WithMessage(err, msg)))
	}
}

func EExt(err error, code int) {
	if err != nil {
		const format = "\x1b[31;1m%s\x1b[0m" //红色打印
		var message = fmt.Sprintf(format, errors.WithMessage(err, "wrong"))
		fmt.Println(message)
		utils_logzap.ZAPS.P1.LOG.Warn("ERROR", zap.Error(errors.WithMessage(err, "wrong")))
		os.Exit(code)
	}
}
