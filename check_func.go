// Package done: Simple error handling package with panic-based validation
// Auto panics on errors to simplify error checking patterns in Go code
// Supports validation functions that reduce boilerplate error handling
// Provides type-safe validation methods with inline assertions
//
// done: Go 错误处理简化包，基于 panic 的验证机制
// 错误时自动 panic，简化 Go 代码中的错误检查模式
// 支持验证函数，减少样板错误处理代码
// 提供类型安全的验证方法和内联断言
package done

import (
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Done panics and logs the stack trace if the error occurs, otherwise do nothing
// Done 当err非空时就panic且打印调用栈信息
func Done(err error) {
	if err != nil {
		zaplog.LOGS.Skip1.Panic("ERROR", zap.Error(err))
	}
}

// Must panics and logs the stack trace if the error occurs, otherwise do nothing
// Must 当err非空时就panic且打印调用栈信息
func Must(err error) {
	if err != nil {
		zaplog.LOGS.Skip1.Panic("ERROR", zap.Error(err))
	}
}

// Soft logs a warning and the stack trace if the error occurs, otherwise do nothing
// Soft 当err非空时就打印warning日志且打印调用栈信息
func Soft(err error) {
	if err != nil {
		zaplog.LOGS.Skip1.Warn("WARN", zap.Error(err))
	}
}

// Fata logs a fatal error and exits the system if the error occurs, otherwise do nothing
// Fata 当出错时就是调用 log.Fatal 退出系统，当不出错时就不做任何事情
func Fata(err error) {
	if err != nil {
		zaplog.LOGS.Skip1.Fatal("ERROR", zap.Error(err))
	}
}
