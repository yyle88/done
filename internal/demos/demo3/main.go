package main

import (
	"fmt"

	"github.com/yyle88/done"
)

func main() {
	// Directly checking the result of run with error handling
	// 使用直接方法检查 run 的结果并处理错误
	demoLogic()

	// Using done.P2 for more concise error handling and result validation
	// 使用 done.P2 使错误处理和结果验证更简洁
	sameLogic()
}

func demoLogic() {
	a, b, err := run()
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

func sameLogic() {
	a, b := done.P2(run()) // done.P2 simplifies error checking and validation
	// done.P2 简化了错误检查和验证
	fmt.Println(a, b) // Output both 'a' and 'b'
	// 输出 'a' 和 'b'
}

// --------------------------
// Simulated business logic
// --------------------------

type A struct {
	ID   int
	Name string
}

type B struct {
	Age  int
	City string
}

func run() (*A, *B, error) {
	a := &A{ID: 1, Name: "Alice"}
	b := &B{Age: 30, City: "New York"}
	return a, b, nil
}
