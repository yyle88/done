package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/yyle88/done"
)

// This example demonstrates two approaches to error handling: manual and compact chained.
// Show how compact invocation simplifies error checking while keeping the logic concise.
// 该示例展示了两种错误处理方式：手动和紧凑链式。
// 示例强调了如何通过紧凑式调用简化错误检查并保持逻辑简洁。

func main() {
	// Manual error handling example
	// 手动处理错误示例
	fmt.Println(ManualErrorHandling())

	// Compact chained error handling example
	// 紧凑链式处理错误示例
	fmt.Println(CompactChainedHandling())
}

// ManualErrorHandling demonstrates step-by-step error checking and handling.
// Pros: Clear and explicit; Cons: Verbose code.
// ManualErrorHandling 展示了逐步的错误检查与处理。
// 优点：清晰且显式；缺点：代码较冗长。
func ManualErrorHandling() int64 {
	snu, err := fetch()
	if err != nil {
		panic(err) // Handle error manually / 手动处理错误
	}
	num, err := toInt(snu)
	if err != nil {
		panic(err)
	}
	if num <= 0 {
		panic(errors.New("num is not > 0"))
	}
	return num
}

// CompactChainedHandling uses `done.VCE` and `done.VNE` for concise error handling.
// Pros: Minimalistic; Cons: Harder to debug if something goes wrong.
// CompactChainedHandling 使用 `done.VCE` 和 `done.VNE` 以简化错误处理。
// 优点：代码极简；缺点：如果出现问题，调试较困难。
func CompactChainedHandling() int64 {
	// Chain methods to fetch, convert, and validate the value
	// 链式调用方法获取、转换并验证值
	return done.VNE(
		toInt(
			done.VCE(fetch()).Nice(),
		),
	).Gt(0)
}

// ----------------------------------------------------------------------
// The following functions simulate business logic for demonstration purposes.
// 以下函数仅为演示用途，模拟了业务逻辑。
// ----------------------------------------------------------------------

// fetch simulates fetching a string value.
// fetch 模拟获取字符串值。
func fetch() (string, error) {
	return "100", nil // Simulated successful fetch / 模拟成功获取
}

// toInt converts a string to an int64 value.
// toInt 将字符串转换为 int64 值。
func toInt(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
