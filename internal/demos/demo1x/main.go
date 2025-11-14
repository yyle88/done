package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/yyle88/done"
)

// This example demonstrates two approaches: step-based and compact chained.
// Shows how compact invocation simplifies validation while keeping the logic concise.
// 该示例展示了两种处理方式：基于步骤和紧凑链式。
// 展示如何通过紧凑式调用简化验证并保持逻辑简洁。

func main() {
	// Classic approach with explicit checks
	// 经典方式的显式检查
	fmt.Println(WithClassicErrorHandling())

	// Compact chained approach
	// 紧凑链式方式
	fmt.Println(WithCompactChainedHandling())
}

// WithClassicErrorHandling demonstrates classic validation with explicit checks.
// Pros: Distinct and explicit; Cons: Verbose code.
// WithClassicErrorHandling 展示了经典的显式检查验证。
// 优点：清晰且显式；缺点：代码较冗长。
func WithClassicErrorHandling() int64 {
	text, err := webFetch()
	if err != nil {
		panic(err) // Handle error at each step // 在每个步骤处理错误
	}
	num, err := parseNum(text)
	if err != nil {
		panic(err)
	}
	if num <= 0 {
		panic(errors.New("num must be positive"))
	}
	return num
}

// WithCompactChainedHandling uses done.VCE and done.VNE to provide concise validation.
// Pros: Compact code; Cons: Debugging can be challenging.
// WithCompactChainedHandling 使用 done.VCE 和 done.VNE 提供简洁的验证。
// 优点：代码紧凑；缺点：调试可能较困难。
func WithCompactChainedHandling() int64 {
	// Chain methods to fetch, parse, and validate the value
	// 链式调用方法获取、解析并验证值
	return done.VNE(
		parseNum(
			done.VCE(webFetch()).Nice(),
		),
	).Gt(0)
}

// ----------------------------------------------------------------------
// The following functions simulate business logic for demonstration.
// 以下函数仅为演示用途，模拟业务逻辑。
// ----------------------------------------------------------------------

// webFetch simulates fetching a string value from a remote source.
// webFetch 模拟从远程源获取字符串值。
func webFetch() (string, error) {
	return "100", nil // Simulated data fetch // 模拟数据获取
}

// parseNum converts a string to an int64 value.
// parseNum 将字符串转换为 int64 值。
func parseNum(text string) (int64, error) {
	return strconv.ParseInt(text, 10, 64)
}
