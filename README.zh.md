[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/done/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/done/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/done)](https://pkg.go.dev/github.com/yyle88/done)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/done/main.svg)](https://coveralls.io/github/yyle88/done?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23%2C%201.24%2C%201.25-lightgrey.svg)](https://github.com/yyle88/done)
[![GitHub Release](https://img.shields.io/github/release/yyle88/done.svg)](https://github.com/yyle88/done/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/done)](https://goreportcard.com/report/github.com/yyle88/done)

# Done - 简化 Go 中的错误处理

**Done** 让你摆脱重复的 `if err != nil` 模式，更专注于业务逻辑。

当你写如下逻辑时：

```go
if err := run(); err != nil {
panic(err)
}
```

使用 Done 简化为：

```go
done.Done(run())
```

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

---

## 特性

- **消除样板代码**：用简洁的函数调用替代冗长的 `if err != nil` 代码块
- **类型安全验证**：基于 Go 泛型构建，提供编译时类型检查和 IDE 智能提示
- **链式操作**：将验证调用链接在一起，代码流畅清晰易读
- **多种返回类型**：处理返回值、指针、切片、map 等各种类型的函数
- **丰富的断言**：专门的验证器覆盖数值、字符串、布尔值和集合
- **聚焦业务逻辑**：错误处理最小化，业务逻辑更突出

---

## 安装

```bash
go get github.com/yyle88/done
```

---

## 使用方法

### 服务链验证

该示例展示了三种方式：经典的逐步验证、链式验证和紧凑链式验证。

```go
package main

import (
	"fmt"

	"github.com/yyle88/done"
)

func main() {
	// 经典方式的显式检查
	fmt.Println(WithClassicErrorHandling())

	// 使用 done.VCE 的链式方式
	fmt.Println(WithChainedErrorHandling())

	// 紧凑链式方式
	fmt.Println(WithCompactChainedHandling())
}

// WithClassicErrorHandling 展示了经典的显式检查验证
// 优点：清晰且简单调试；缺点：代码冗长
func WithClassicErrorHandling() string {
	service, err := NewService()
	if err != nil {
		panic(err) // 在每个步骤处理错误
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

// WithChainedErrorHandling 使用 done.VCE 以链式形式简化验证
// 优点：代码紧凑；缺点：调试堆栈可能较困难
func WithChainedErrorHandling() string {
	service := done.VCE(NewService()).Nice()
	client := done.VCE(service.GetClient()).Nice()
	response := done.VCE(client.GetResponse()).Nice()
	return response.Message
}

// WithCompactChainedHandling 展示最紧凑的链式验证形式
// 优点：代码简洁；缺点：调试较困难
func WithCompactChainedHandling() string {
	return done.VCE(done.VCE(done.VCE(
		NewService(),
	).Nice().GetClient(),
	).Nice().GetResponse(),
	).Nice().Message
}

// Service 表示链中的核心服务
type Service struct{}

// NewService 创建一个新的 Service 实例
func NewService() (*Service, error) {
	return &Service{}, nil
}

// GetClient 从该服务返回一个 Client 实例
func (s *Service) GetClient() (*Client, error) {
	return &Client{}, nil
}

// Client 表示链中的中间客户端
type Client struct{}

// GetResponse 返回包含结果消息的 Response
func (c *Client) GetResponse() (*Response, error) {
	return &Response{
		Message: "success", // 模拟成功消息
	}, nil
}

// Response 表示包含结果的响应
type Response struct {
	Message string // 结果消息
}
```

⬆️ **源码:** [源码](internal/demos/demo3x/main.go)

### 链式操作

该示例展示了两种错误处理方式：经典的逐步验证和使用 done 的紧凑链式验证。

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/yyle88/done"
)

func main() {
	// 经典方式的显式检查
	fmt.Println(WithClassicErrorHandling())

	// 紧凑链式方式
	fmt.Println(WithCompactChainedHandling())
}

// WithClassicErrorHandling 展示了经典的显式检查验证
// 优点：清晰且显式；缺点：代码较冗长
func WithClassicErrorHandling() int64 {
	text, err := webFetch()
	if err != nil {
		panic(err) // 在每个步骤处理错误
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

// WithCompactChainedHandling 使用 done.VCE 和 done.VNE 提供简洁的验证
// 优点：代码紧凑；缺点：调试可能较困难
func WithCompactChainedHandling() int64 {
	// 链式调用方法获取、解析并验证值
	return done.VNE(
		parseNum(
			done.VCE(webFetch()).Nice(),
		),
	).Gt(0)
}

// webFetch 模拟从远程源获取字符串值
func webFetch() (string, error) {
	return "100", nil // 模拟数据获取
}

// parseNum 将字符串转换为 int64 值
func parseNum(text string) (int64, error) {
	return strconv.ParseInt(text, 10, 64)
}
```

⬆️ **源码:** [源码](internal/demos/demo1x/main.go)

### 指针验证

该示例展示了使用 done.P2 验证和提取多个指针结果。

```go
package main

import (
	"fmt"

	"github.com/yyle88/done"
)

func main() {
	// 经典方式：使用显式验证检查结果
	WithClassicErrorHandling()

	// 紧凑方式：使用 done.P2 减少样板代码
	WithCompactErrorHandling()
}

// WithClassicErrorHandling 展示经典的显式检查验证
// 优点：显式且清晰；缺点：代码冗长
func WithClassicErrorHandling() {
	account, config, err := fetchAccountAndConfig()
	if err != nil {
		panic(err) // 检查每个步骤
	}
	if account == nil {
		panic("account is nil") // 验证账户存在
	}
	if config == nil {
		panic("config is nil") // 验证配置存在
	}
	fmt.Println(account, config) // 打印账户和配置
}

// WithCompactErrorHandling 使用 done.P2 精简验证
// 优点：代码简洁；缺点：验证是隐式的
func WithCompactErrorHandling() {
	account, config := done.P2(fetchAccountAndConfig()) // done.P2 处理检查
	fmt.Println(account, config)                        // 打印账户和配置
}

// Account 表示系统中的账户
type Account struct {
	ID   int    // 账户ID
	Name string // 账户名称
}

// Config 表示配置设置
type Config struct {
	Timeout int    // 超时时间（秒）
	Region  string // 服务区域
}

// fetchAccountAndConfig 模拟获取账户和配置数据
func fetchAccountAndConfig() (*Account, *Config, error) {
	account := &Account{ID: 1, Name: "Alice"}
	config := &Config{Timeout: 30, Region: "us-west"}
	return account, config, nil
}
```

⬆️ **源码:** [源码](internal/demos/demo2x/main.go)

---

## 核心类型

| 类型                      | 描述                                        |
|-------------------------|-------------------------------------------|
| **`Ve[V any]`**         | 封装任意值和错误，提供 `Done`、`Must`、`Soft` 来处理错误 |
| **`Vpe[V any]`**        | 封装指针和错误，使用 `Sure`、`Nice`、`Full` 验证非空指针 |
| **`Vce[V comparable]`** | 封装可比较值和错误，使用 `Same`、`Diff`、`Equals` 比较值 |
| **`Vbe`**               | 封装布尔值和错误，使用 `TRUE`、`FALSE`、`OK`、`NO` 断言真假 |
| **`Vae[V any]`**        | 封装切片和错误，使用 `Some`、`Have`、`Length` 检查空和长度 |
| **`Vme[K, V]`**         | 封装 map 和错误，使用 `Nice`、`Size`、`Len` 验证大小和内容 |

---

## 核心函数

| 函数         | 描述                 |
|------------|--------------------|
| **`Done`** | 错误存在时记录日志并触发 panic |
| **`Must`** | 确保错误为 nil，成功时返回值 |
| **`Soft`** | 记录告警不触发 panic，继续执行 |
| **`Fata`** | 记录致命错误并终止程序 |

---

## 功能分组

| 分组       | 函数                     | 描述                    |
|----------|------------------------|-----------------------|
| **错误** | `Done`, `Must`, `Soft` | 根据严重阶段使用 panic 或告警处理错误 |
| **非零验证** | `Sure`, `Nice`, `Some` | 验证非零值并返回 |
| **非零断言** | `Good`, `Fine`, `Safe` | 断言非零值但不返回 |
| **零值检查** | `Zero`, `None`, `Void` | 验证值为零值或空 |
| **值比较** | `Same`, `Diff`, `Is`, `Equals` | 比较值检查匹配或差异 |

---

## 类型特定验证

| 类型        | 用途 | 方法 |
|-----------|------|------|
| **`Vce`** | 可比较值 | `Same`、`Diff`、`Is`、`Equals` - 值比较 |
| **`Vse`** | 字符串操作 | `HasPrefix`、`HasSuffix`、`Contains` - 子串检查 |
| **`Vne`** | 数值比较 | `Gt`、`Lt`、`Gte`、`Lte` - 范围验证 |
| **`Vbe`** | 布尔断言 | `TRUE`、`FALSE`、`YES`、`NO`、`OK` - 真假检查 |
| **`Vae`** | 切片验证 | `Sure`、`Some`、`Have`、`Length` - 空和大小检查 |
| **`Vme`** | Map 验证 | `Nice`、`Some`、`Size`、`Len` - map 大小和内容检查 |

---

## 示例

### 基础错误处理

**简单错误检查：**
```go
done.Done(run())
```

**返回值与错误检查：**
```go
result := done.V1(fetchData())
```

**多个返回值：**
```go
v1, v2 := done.V2(getTwoValues())
```

### 指针验证

**单指针验证：**
```go
ptr := done.P1(getPointer())
```

**多指针验证：**
```go
ptr1, ptr2 := done.P2(getTwoPointers())
```

### 可比较值操作

**检查值相同：**
```go
value := done.VCE(getValue()).Same(expected)
```

**检查值不同：**
```go
value := done.VCE(getValue()).Diff(unwanted)
```

### 数值比较

**验证值超过阈值：**
```go
num := done.VNE(getNumber()).Gt(0)
```

**小于验证：**
```go
num := done.VNE(getNumber()).Lt(100)
```

### 布尔值验证

**确保布尔值为真：**
```go
done.VBE(checkCondition()).TRUE()
```

**确保布尔值为假：**
```go
done.VBE(checkCondition()).FALSE()
```

### 切片操作

**确保切片有元素：**
```go
items := done.VAE(getSlice()).Some()
```

**检查切片长度：**
```go
items := done.VAE(getSlice()).Length(3)
```

### Map 操作

**确保 map 有内容：**
```go
data := done.VME(getMap()).Nice()
```

**检查 map 大小：**
```go
data := done.VME(getMap()).Size(5)
```

---

## 总结

**Done** 包为 Go 语言带来可靠的错误处理方式。借助消除重复的错误检查，让你编写简洁、易维护的代码。在构建原型和生产系统时，Done 都能帮助你聚焦于最重要的事情：业务逻辑。

**测试一下，告诉我们你的想法！**

---

## 关联项目

探索此生态系统中更多错误处理包：

### 高级包

- **[must](https://github.com/yyle88/must)** - Must 风格断言，丰富的类型支持和详细的错误上下文
- **[rese](https://github.com/yyle88/rese)** - 结果提取与 panic，专注于安全解包值

### 基础包

- **[done](https://github.com/yyle88/done)** - 简单专注的错误处理（本项目）
- **[sure](https://github.com/yyle88/sure)** - 生成代码来创建自定义验证方法

每个包针对不同的使用场景，从快速原型到具有全面错误处理的生产系统。

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->
