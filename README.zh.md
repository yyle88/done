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

## 英文文档

[English README](README.md)

---

## 特性

- **简单的错误处理**：快速处理错误，使用简单的函数。
- **支持多个返回值**：适用于返回值和错误类型的函数。
- **减少重复的代码**：用简洁明了的函数代替重复的错误检查。

---

## 安装

```bash
go get github.com/yyle88/done
```

---

## 工作原理

**Done** 包将错误检查封装为易用的断言工具，帮助你专注于代码的成功路径，而不是错误处理的细节。

---

## 核心类型

| 类型                      | 描述                                        |
|-------------------------|-------------------------------------------|
| **`Ve[V any]`**         | 包含一个值和一个错误，提供 `Done`、`Must` 和 `Soft` 等操作。 |
| **`Vpe[V any]`**        | 处理指针值，包含 `Sure`、`Nice` 和 `Good` 等操作。      |
| **`Vce[V comparable]`** | 用于可比较值，提供 `Same`、`Diff` 和 `Equals` 等操作。   |

---

## 核心函数

| 函数         | 描述                 |
|------------|--------------------|
| **`Done`** | 如果存在错误则触发 panic。   |
| **`Must`** | 确保错误为 nil 并返回值。    |
| **`Soft`** | 记录警告错误，但不触发 panic。 |
| **`Fata`** | 记录错误日志并触发 panic。   |

---

## 功能概览

| 类别       | 函数                     | 描述                    |
|----------|------------------------|-----------------------|
| **错误处理** | `Done`, `Must`, `Soft` | 对错误进行 panic 或特定方式的处理。 |
| **结果验证** | `Sure`, `Nice`, `Some` | 确保结果非零值并返回。           |
| **结果验证** | `Good`, `Fine`, `Safe` | 确保结果非零值但不返回结果。        |
| **零值检查** | `Zero`, `None`, `Void` | 确保结果为其类型的零值。          |
| **量值比较** | `Same`, `Diff`, `Is`   | 检查值是否相同、不同或符合特定条件。    |

---

## 高级错误处理

| 工具        | 描述                                                   |
|-----------|------------------------------------------------------|
| **`Vce`** | 针对可比较值，提供 `Same`、`Diff`、`Is` 和 `Equals` 等方法。         |
| **`Vse`** | 针对字符串操作，提供 `HasPrefix`、`HasSuffix` 和 `Contains` 等方法。 |
| **`Vne`** | 针对数字操作，提供 `Gt`（大于）、`Lt`（小于）等方法。                      |

---

## 示例用法

### 标准错误处理

```go
package main

import (
	"fmt"
)

func main() {
	xyz, err := NewXyz()
	if err != nil {
		panic(err) // 手动处理错误
	}
	abc, err := xyz.Abc()
	if err != nil {
		panic(err)
	}
	uvw, err := abc.Uvw()
	if err != nil {
		panic(err)
	}
	fmt.Println(uvw.Message)
}
```

### 使用 Done

```go
package main

import (
	"fmt"
	"github.com/yyle88/done"
)

func main() {
	xyz := done.VCE(NewXyz()).Nice()
	abc := done.VCE(xyz.Abc()).Nice()
	uvw := done.VCE(abc.Uvw()).Nice()
	fmt.Println(uvw.Message)
}
```

这种方法通过链式调用简化了代码，同时结合错误处理断言，使代码更简洁、易读。

---

## 链式操作

### 不使用 Done

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

func main() {
	stringNum, err := fetch()
	if err != nil {
		panic(err)
	}
	num, err := toInt(stringNum)
	if err != nil {
		panic(err)
	}
	if num <= 0 {
		panic(errors.New("num <= 0"))
	}
	fmt.Println(num)
}
```

### 使用 Done

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/yyle88/done"
)

func main() {
	num := done.VNE(toInt(done.VCE(fetch()).Nice())).Gt(0)
	fmt.Println(num)
}
```

通过 **Done**，你可以消除重复的错误检查，用内联断言实现条件校验，使代码更整洁、可维护。

---

## 总结

**Done** 包是一个强大的工具，可以简化错误处理，特别适用于非正式或小型项目。它减少了模板代码的重复，使错误处理更简洁，让你专注于编写清晰高效的业务逻辑。

**试试看并分享你的反馈吧！**

---

## 许可证

此项目基于 MIT 许可证。详细信息请参阅 [LICENSE](LICENSE) 文件。

---

## 贡献与支持

欢迎通过提交拉取请求或报告问题来为此项目做出贡献。

如果你觉得这个包有用，请在 GitHub 上给个星！

**感谢你的支持！**
