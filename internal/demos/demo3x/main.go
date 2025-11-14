package main

import (
	"fmt"

	"github.com/yyle88/done"
)

// This example demonstrates three approaches: step-based, chained, and compact.
// The focus is showing how compact invocation can reduce boilerplate in appropriate contexts.
// 该示例展示了三种处理方式：基于步骤、链式和紧凑式
// 重点展示如何在合适的场景下通过紧凑式调用减少样板代码

func main() {
	// Classic approach with explicit checks
	// 经典方式的显式检查
	fmt.Println(WithClassicErrorHandling())

	// Chained approach with done.VCE
	// 使用 done.VCE 的链式方式
	fmt.Println(WithChainedErrorHandling())

	// Compact chained approach
	// 紧凑链式方式
	fmt.Println(WithCompactChainedHandling())
}

// WithClassicErrorHandling demonstrates classic validation with explicit checks.
// Pros: Distinct and simple to debug; Cons: Verbose code.
// WithClassicErrorHandling 展示了经典的显式检查验证。
// 优点：清晰且简单调试；缺点：代码冗长。
func WithClassicErrorHandling() string {
	service, err := NewService()
	if err != nil {
		panic(err) // Handle errors at each step // 在每个步骤处理错误
	}
	client, err := service.GetClient()
	if err != nil {
		panic(err)
	}
	response, err := client.GetResponse()
	if err != nil {
		panic(err)
	}
	return response.Message
}

// WithChainedErrorHandling uses done.VCE to streamline validation in a chained fashion.
// Pros: Compact code; Cons: Debugging stacks can be challenging.
// WithChainedErrorHandling 使用 done.VCE 以链式形式简化验证。
// 优点：代码紧凑；缺点：调试堆栈可能较困难。
func WithChainedErrorHandling() string {
	service := done.VCE(NewService()).Nice()
	client := done.VCE(service.GetClient()).Nice()
	response := done.VCE(client.GetResponse()).Nice()
	return response.Message
}

// WithCompactChainedHandling shows the most compact form of chained validation.
// Pros: Concise code; Cons: Debugging is challenging.
// WithCompactChainedHandling 展示最紧凑的链式验证形式。
// 优点：代码简洁；缺点：调试较困难。
func WithCompactChainedHandling() string {
	return done.VCE(done.VCE(done.VCE(
		NewService(),
	).Nice().GetClient(),
	).Nice().GetResponse(),
	).Nice().Message
}

// ----------------------------------------------------------------------
// Simulated business logic structures and methods for demonstration.
// 以下部分仅为演示用途，模拟业务逻辑。
// ----------------------------------------------------------------------

// Service represents the primary service in the chain.
// Service 表示链中的主要服务。
type Service struct{}

// NewService creates a new Service instance.
// NewService 创建一个新的 Service 实例。
func NewService() (*Service, error) {
	return &Service{}, nil
}

// GetClient returns a Client instance from this service.
// GetClient 从该服务返回一个 Client 实例。
func (s *Service) GetClient() (*Client, error) {
	return &Client{}, nil
}

// Client represents the intermediate client in the chain.
// Client 表示链中的中间客户端。
type Client struct{}

// GetResponse returns the Response containing the result message.
// GetResponse 返回包含结果消息的 Response。
func (c *Client) GetResponse() (*Response, error) {
	return &Response{
		Message: "success", // Simulated success message // 模拟成功消息
	}, nil
}

// Response represents the response containing the result.
// Response 表示包含结果的响应。
type Response struct {
	Message string // Result message // 结果消息
}
