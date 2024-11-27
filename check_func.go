package done

import (
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Done 当err非空是就panic且打印调用栈信息
func Done(err error) {
	if err != nil {
		zaplog.LOGS.P1.Panic("ERROR", zap.Error(err))
	}
}

// Must 当err非空是就panic且打印调用栈信息
func Must(err error) {
	if err != nil {
		zaplog.LOGS.P1.Panic("ERROR", zap.Error(err))
	}
}

// Soft 当err非空是就打印warning日志且打印调用栈信息
func Soft(err error) {
	if err != nil {
		zaplog.LOGS.P1.Warn("WARN", zap.Error(err))
	}
}

// Fata 当出错时就是调用 log.Fatal 退出系统，当不出错时就不做任何事情
// FaFaFaFaFa 哈哈哈哈哈
func Fata(err error) {
	if err != nil {
		zaplog.LOGS.P1.Fatal("ERROR", zap.Error(err))
	}
}
