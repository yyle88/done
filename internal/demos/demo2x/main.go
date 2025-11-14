package main

import (
	"fmt"

	"github.com/yyle88/done"
)

func main() {
	// Classic approach: check results with explicit validation
	// 经典方式：使用显式验证检查结果
	WithClassicErrorHandling()

	// Compact approach: use done.P2 to reduce boilerplate
	// 紧凑方式：使用 done.P2 减少样板代码
	WithCompactErrorHandling()
}

// WithClassicErrorHandling demonstrates classic validation with explicit checks.
// Pros: Explicit and distinct; Cons: Verbose code.
// WithClassicErrorHandling 展示经典的显式检查验证。
// 优点：显式且清晰；缺点：代码冗长。
func WithClassicErrorHandling() {
	account, config, err := fetchAccountAndConfig()
	if err != nil {
		panic(err) // Check each step // 检查每个步骤
	}
	if account == nil {
		panic("account is nil") // Validate account exists // 验证账户存在
	}
	if config == nil {
		panic("config is nil") // Validate config exists // 验证配置存在
	}
	fmt.Println(account, config) // Print both account and config // 打印账户和配置
}

// WithCompactErrorHandling uses done.P2 to simplify validation.
// Pros: Concise code; Cons: Validation is implicit.
// WithCompactErrorHandling 使用 done.P2 简化验证。
// 优点：代码简洁；缺点：验证是隐式的。
func WithCompactErrorHandling() {
	account, config := done.P2(fetchAccountAndConfig()) // done.P2 handles checks // done.P2 处理检查
	fmt.Println(account, config)                        // Print both account and config // 打印账户和配置
}

// ----------------------------------------------------------------------
// Simulated business logic with demonstration data.
// 以下部分仅为演示用途，模拟业务逻辑。
// ----------------------------------------------------------------------

// Account represents an account in the system.
// Account 表示系统中的账户。
type Account struct {
	ID   int    // Account ID // 账户ID
	Name string // Account name // 账户名称
}

// Config represents configuration settings.
// Config 表示配置设置。
type Config struct {
	Timeout int    // Timeout in seconds // 超时时间（秒）
	Region  string // Service region // 服务区域
}

// fetchAccountAndConfig simulates fetching account and config data.
// fetchAccountAndConfig 模拟获取账户和配置数据。
func fetchAccountAndConfig() (*Account, *Config, error) {
	account := &Account{ID: 1, Name: "Alice"}
	config := &Config{Timeout: 30, Region: "us-west"}
	return account, config, nil
}
