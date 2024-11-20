package main

import (
	"fmt"

	"github.com/yyle88/done"
)

func main() {
	// Directly checking the result of run1 with error handling
	// 使用直接方法检查 run1 的结果并处理错误
	{
		a, err := run1()
		if err != nil {
			panic(err) // If an error occurs, trigger a panic
			// 如果发生错误，触发 panic
		}
		if a == nil {
			panic("a is nil") // If 'a' is nil, trigger panic
			// 如果 'a' 为 nil，触发 panic
		}
		fmt.Println(a) // Output the result
		// 输出结果
	}

	// Using done.P1 for cleaner error handling and result validation
	// 使用 done.P1 简化错误处理和结果验证
	{
		a := done.P1(run1()) // done.P1 handles error checking and nil validation
		// done.P1 处理错误检查和 nil 验证
		fmt.Println(a) // Output the result
		// 输出结果
	}

	// Directly checking the result of run2 with error handling
	// 使用直接方法检查 run2 的结果并处理错误
	{
		a, b, err := run2()
		if err != nil {
			panic(err) // If an error occurs, trigger a panic
			// 如果发生错误，触发 panic
		}
		if a == nil {
			panic("a is nil") // If 'a' is nil, trigger panic
			// 如果 'a' 为 nil，触发 panic
		}
		if b == nil {
			panic("b is nil") // If 'b' is nil, trigger panic
			// 如果 'b' 为 nil，触发 panic
		}
		fmt.Println(a, b) // Output both 'a' and 'b'
		// 输出 'a' 和 'b'
	}

	// Using done.P2 for more concise error handling and result validation
	// 使用 done.P2 使错误处理和结果验证更简洁
	{
		a, b := done.P2(run2()) // done.P2 simplifies error checking and validation
		// done.P2 简化了错误检查和验证
		fmt.Println(a, b) // Output both 'a' and 'b'
		// 输出 'a' 和 'b'
	}

	// Directly checking the result of run3 with error handling
	// 使用直接方法检查 run3 的结果并处理错误
	{
		a, b, c, err := run3()
		if err != nil {
			panic(err) // If an error occurs, trigger a panic
			// 如果发生错误，触发 panic
		}
		if a == nil {
			panic("a is nil") // If 'a' is nil, trigger panic
			// 如果 'a' 为 nil，触发 panic
		}
		if b == nil {
			panic("b is nil") // If 'b' is nil, trigger panic
			// 如果 'b' 为 nil，触发 panic
		}
		if c == nil {
			panic("c is nil") // If 'c' is nil, trigger panic
			// 如果 'c' 为 nil，触发 panic
		}
		fmt.Println(a, b, c) // Output 'a', 'b', and 'c'
		// 输出 'a'，'b' 和 'c'
	}

	// Using done.P3 to simplify error handling and result validation for run3
	// 使用 done.P3 简化 run3 的错误处理和结果验证
	{
		a, b, c := done.P3(run3()) // done.P3 simplifies error checking and validation
		// done.P3 简化了错误检查和验证
		fmt.Println(a, b, c) // Output 'a', 'b', and 'c'
		// 输出 'a'，'b' 和 'c'
	}
}

// --------------------------
// Simulated business logic
// --------------------------

// Defining mock types for demonstration with fields
// 为演示定义模拟类型并添加字段
type A struct {
	ID   int
	Name string
}

type B struct {
	Age  int
	City string
}

type C struct {
	Price float64
	Code  string
}

// run1 simulates returning an instance of A and error
// run1 模拟返回 A 的实例和错误
func run1() (*A, error) {
	return &A{ID: 1, Name: "Alice"}, nil
}

// run2 simulates returning instances of A and B, and error
// run2 模拟返回 A 和 B 的实例以及错误
func run2() (*A, *B, error) {
	return &A{ID: 1, Name: "Alice"}, &B{Age: 30, City: "New York"}, nil
}

// run3 simulates returning instances of A, B, and C, and error
// run3 模拟返回 A, B 和 C 的实例以及错误
func run3() (*A, *B, *C, error) {
	return &A{ID: 1, Name: "Alice"}, &B{Age: 30, City: "New York"}, &C{Price: 100.5, Code: "XYZ123"}, nil
}
