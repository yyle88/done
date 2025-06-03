[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/done/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/done/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/done)](https://pkg.go.dev/github.com/yyle88/done)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/done/master.svg)](https://coveralls.io/github/yyle88/done?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/done.svg)](https://github.com/yyle88/done/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/done)](https://goreportcard.com/report/github.com/yyle88/done)

# Done - Simple Error Handling in Go

**Done** allows you to focus on your business logic without repetitive `if err != nil` patterns.

when you write logic:

```
if err := run(); err != nil{
    panic(err)
}
```

then you can use done:

```
done.Done(run())
```

## CHINESE README

[中文说明](README.zh.md)

---

## Features

- **Simple error handling**: Handle errors quickly with simple functions.
- **Supports multiple results**: Works well with functions that return a value and an error type.
- **Less repeated code**: Replace repeated error checks with short and clear functions.

---

## Installation

```bash
go get github.com/yyle88/done
```

---

## How It Works

The **Done** package wraps error checking into easy-to-use assertions, enabling you to handle errors while focusing on the success path of code.

---

## Core Types

| Type                    | Description                                                                   |
|-------------------------|-------------------------------------------------------------------------------|
| **`Ve[V any]`**         | Holds a value and an error. Provides methods like `Done`, `Must`, and `Soft`. |
| **`Vpe[V any]`**        | For pointer values, includes methods such as `Sure`, `Nice`, and `Good`.      |
| **`Vce[V comparable]`** | For comparable values, provides methods like `Same`, `Diff`, and `Equals`.    |

---

## Key Functions

| Function   | Description                                       |
|------------|---------------------------------------------------|
| **`Done`** | Panics if an error exists.                        |
| **`Must`** | Ensures the error is `nil` and returns the value. |
| **`Soft`** | Logs a warning for errors but does not panic.     |
| **`Fata`** | Logs the error at a "fatal" level and panics.     |

---

## Functionality Summary

| Category                | Functions                      | Description                                                             |
|-------------------------|--------------------------------|-------------------------------------------------------------------------|
| **Error Handling**      | `Done`, `Must`, `Soft`         | Panics on error or handles it in a specific way.                        |
| **Result Validation**   | `Sure`, `Nice`, `Some`         | Ensures the result is not zero and returns it.                          |
| **Result Validation**   | `Good`, `Fine`, `Safe`         | Ensures the result is not zero without returning it.                    |
| **Zero Value Checking** | `Zero`, `None`, `Void`         | Ensures the result is the zero value of its type.                       |
| **Value Comparisons**   | `Same`, `Diff`, `Is`, `Equals` | Checks if values are the same, different, or match specific conditions. |

---

## Advanced Error Handling

| Utility   | Description                                                                               |
|-----------|-------------------------------------------------------------------------------------------|
| **`Vce`** | For comparable values, includes methods like `Same`, `Diff`, `Is`, and `Equals`.          |
| **`Vse`** | For string operations, includes methods like `HasPrefix`, `HasSuffix`, and `Contains`.    |
| **`Vne`** | For numeric operations, includes methods like `Gt` (greater than), `Lt` (less than), etc. |

---

## Example Usage

### Standard Error Handling

```go
package main

import (
	"fmt"
)

func main() {
	xyz, err := NewXyz()
	if err != nil {
		panic(err) // Handle errors manually
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

### Using Done

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

This approach simplifies the code by chaining function calls with error handling assertions.

---

## Chaining Operations

### Without Done

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

### With Done

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

By using **Done**, you eliminate repetitive error checks and enable inline assertions for conditions, resulting in cleaner and more maintainable code.

---

## Conclusion

The **Done** package is a powerful tool for simplifying error handling, especially in informal or small-scale projects. It helps reduce boilerplate, makes error handling concise, and lets you focus on writing clean and efficient business logic.

**Give it a try and share your feedback!**

---

## License

MIT License. See [LICENSE](LICENSE).

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repo on GitHub (using the webpage interface).
2. Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. Navigate to the cloned project (`cd repo-name`)
4. Create a feature branch (`git checkout -b feature/xxx`).
5. Stage changes (`git add .`)
6. Commit changes (`git commit -m "Add feature xxx"`).
7. Push to the branch (`git push origin feature/xxx`).
8. Open a pull request on GitHub (on the GitHub webpage).

Please ensure tests pass and include relevant documentation updates.

---

## Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

If you find this package valuable, give me some stars on GitHub! Thank you!!!

**Thank you for your support!**

**Happy Coding with this package!** 🎉

Give me stars. Thank you!!!

---

## GitHub Stars

[![starring](https://starchart.cc/yyle88/done.svg?variant=adaptive)](https://starchart.cc/yyle88/done)
