package main

import (
	"fmt"

	"github.com/yyle88/done"
)

// This example demonstrates three approaches to error handling: manual, chained, and compact chained.
// The primary focus is showcasing how compact invocation can simplify error handling in appropriate contexts.
// 该示例展示了三种错误处理方式：手动、链式和紧凑链式，主要是演示如何在合适的情况下通过紧凑式调用简化错误处理。

func main() {
	// Manual error handling
	// 手动处理错误
	fmt.Println(ManualErrorHandling())

	// Chained error handling
	// 链式处理错误
	fmt.Println(ChainedErrorHandling())

	// Compact chained error handling
	// 紧凑链式处理错误
	fmt.Println(CompactChainedHandling())
}

// ManualErrorHandling demonstrates step-by-step manual error handling.
// Pros: Clear and easy to debug; Cons: Verbose code.
// ManualErrorHandling 展示了逐步手动处理每一步错误的方式。
// 优点：清晰且便于调试；缺点：代码冗长。
func ManualErrorHandling() string {
	xyz, err := NewXyz()
	if err != nil {
		panic(err) // Handle errors manually / 手动处理错误
	}
	abc, err := xyz.Abc()
	if err != nil {
		panic(err)
	}
	uvw, err := abc.Uvw()
	if err != nil {
		panic(err)
	}
	return uvw.Message
}

// ChainedErrorHandling uses `done.VCE` to streamline error handling in a chained manner.
// Pros: Cleaner code; Cons: Debugging error stacks can be challenging.
// ChainedErrorHandling 使用 `done.VCE` 以链式方式简化错误处理。
// 优点：代码更简洁；缺点：调试错误堆栈可能更困难。
func ChainedErrorHandling() string {
	xyz := done.VCE(NewXyz()).Nice()
	abc := done.VCE(xyz.Abc()).Nice()
	uvw := done.VCE(abc.Uvw()).Nice()
	return uvw.Message
}

// CompactChainedHandling represents the most compact form of chained error handling.
// Pros: Extremely concise; Cons: Error debugging is difficult.
// CompactChainedHandling 使用最紧凑的链式写法。
// 优点：代码极简；缺点：定位问题较难。
func CompactChainedHandling() string {
	return done.VCE(done.VCE(done.VCE(
		NewXyz(),
	).Nice().Abc(),
	).Nice().Uvw(),
	).Nice().Message
}

// ----------------------------------------------------------------------
// Simulated business logic structures and methods for demonstration purposes.
// 以下部分仅为演示用途，模拟业务逻辑。
// ----------------------------------------------------------------------

// XyzType represents the first layer of the object hierarchy.
// XyzType 表示对象层次的第一层。
type XyzType struct{}

// NewXyz creates a new XyzType instance.
// NewXyz 创建一个新的 XyzType 实例。
func NewXyz() (*XyzType, error) {
	return &XyzType{}, nil
}

// Abc returns the second-layer object, AbcType.
// Abc 返回第二层对象 AbcType。
func (xyz *XyzType) Abc() (*AbcType, error) {
	return &AbcType{}, nil
}

// AbcType represents the second layer of the object hierarchy.
// AbcType 表示对象层次的第二层。
type AbcType struct{}

// Uvw returns the final object, UvwType, containing the result message.
// Uvw 返回最终的结果对象 UvwType，包含结果消息。
func (abc *AbcType) Uvw() (*UvwType, error) {
	return &UvwType{
		Message: "success", // Simulated success message / 模拟成功消息
	}, nil
}

// UvwType represents the final object containing the result.
// UvwType 表示包含结果的最终对象。
type UvwType struct {
	Message string
}
