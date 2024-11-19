# Done - Error Handling Simplified in Go

**Done** is a utility package designed to simplify error handling in Go projects, especially in informal or small-scale applications. It provides a set of functions and types to reduce the verbosity of error checking and handling by using assertions for successful operations. This package can help you avoid repetitive `if err != nil` checks and allow your code to be cleaner, more readable, and faster to write.

## README
[中文说明](README.zh.md)

## Features

- **Simplified error handling:** Instead of explicitly checking for errors at every step, you can use assertions to handle errors automatically.
- **Supports multiple return values:** Effectively handles functions that return both a result and an error.
- **Reduces boilerplate code:** By using short, meaningful function calls, you can skip error checks and focus on your business logic.

## Installation

```bash
go get github.com/yyle88/done
```

## How It Works

This package introduces several types and functions that make error handling more concise. The core idea is to wrap the error check in an assertion, allowing you to proceed with the result if no error occurs.

### Core Types

- `Ve[V any]`: A type that holds a value and an error. Provides methods to handle errors like `Done`, `Must`, and `Soft`.
- `Vpe[V any]`: A type for pointer values, with extended methods such as `Sure`, `Nice`, `Good`, and more.
- `Vce[V comparable]`: A type for comparable values that provides methods like `Same`, `Diff`, `Is`, `Equals`, etc., to handle value comparisons in error-prone scenarios.

### Key Functions

- `Done`: Checks if there is an error, panicking if one exists.
- `Must`: Similar to `Done`, but is more expressive of its intent to assert the error is `nil` and return the value.
- `Soft`: Logs a warning if there is an error but does not panic.
- `Fata`: Panics and logs the error at the "Fatal" level.
- `EExt`: Exits the program with a specific exit code if an error occurs, after printing the error in RED.

### Example Usage

Here is an example code:

```go
package main

import (
    "fmt"
    "github.com/yyle88/done"
)

func main() {
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
	fmt.Println(uvw.Message)
}
```

Here is how you would use the package to handle errors in a typical Go application:

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

See [demo1](internal/demos/demo1/main.go)

In this example, the errors are checked and handled inline, allowing you to chain function calls without cluttering the code with repetitive error checks.

### Chaining Operations

Here is an example code:

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/yyle88/done"
)

func main() {
	snu, err := fetch()
	if err != nil {
		panic(err) // Handle error manually / 手动处理错误
	}
	num, err := toInt(snu)
	if err != nil {
		panic(err)
	}
	if !(num > 0) {
		panic(errors.New("num is not > 0"))
	}
	fmt.Println(num)
}
```

You can chain multiple function calls with error handling assertions in a single line:

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/yyle88/done"
)

func main() {
	num := done.VNE(   toInt(   done.VCE(fetch()).Nice()   )   ).Gt(0)
	fmt.Println(num)
}
```

See [demo2](internal/demos/demo2/main.go)

This way, you can check if the error exists at each step and continue with the result without writing long error-handling logic.

## Usage Scenarios

This package is ideal for projects where:

- You want a concise error handling pattern.
- You are working in small, informal projects where errors are highly unlikely.
- You are okay with using panics for error handling in places where failure is considered unexpected and non-recoverable.

**Important:** Although this package simplifies error handling, be mindful of using panics in production environments. It's suitable for scenarios where failure is extremely rare, such as development, testing, or specific low-risk situations.

## Functionality Summary

The utility functions provided by **Done** are based on different approaches to error handling:

- **Done**: Only checks if `err != nil` and panics.
- **Sure/Nice/Some**: Ensures that the result is not zero (i.e., not `0`, `false`, `nil`, or an empty slice).
- **Good/Fine/Safe**: Verifies that the result is not zero but only checks it (without returning the result).
- **Zero/None/Void**: Ensures that the result is zero (i.e., the zero value of its type).
- **Same/Diff**: Ensures the result is the same or different as the expected value.

## Common Functions

- `Done`: Checks and panics if the error exists.
- `Sure/Nice/Some`: Ensures the result is not zero and returns it.
- `Good/Fine/Safe`: Ensures the result is not zero without returning it.
- `Zero/None/Void`: Ensures the result is zero and returns it.

## Advanced Error Handling

In cases where you need more complex logic, such as comparing values or checking conditions, you can use the following functions:

- **Vce (Comparable + Errors)**: Check if two comparable values are the same or different using `Same`, `Diff`, `Is`, and `Equals`.
- **Vse (String + Errors)**: Check if a string matches conditions like `HasPrefix`, `HasSuffix`, or `Contains`.
- **Vne (Numerical + Errors)**: Check if a numeric value meets conditions like `Gt`, `Lt`, `Gte`, or `Lte`.

## Conclusion

The **Done** package is a powerful and concise utility for error handling, especially in informal projects or small-scale applications. It helps you reduce boilerplate code, simplify error checks, and focus on the core logic of your application.

**Give it a try, and let me know what you think!**

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Give stars

Feel free to contribute or improve the package! Your stars and pull requests are welcome.

## Thank You

If you find this package valuable, give it a star on GitHub! Thank you!!!
