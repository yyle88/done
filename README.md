[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/done/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/done/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/done)](https://pkg.go.dev/github.com/yyle88/done)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/done/main.svg)](https://coveralls.io/github/yyle88/done?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23%2C%201.24%2C%201.25-lightgrey.svg)](https://github.com/yyle88/done)
[![GitHub Release](https://img.shields.io/github/release/yyle88/done.svg)](https://github.com/yyle88/done/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/done)](https://goreportcard.com/report/github.com/yyle88/done)

# Done - Simple Error Handling in Go

**Done** lets you focus on the business logic without repetitive `if err != nil` patterns.

When you write logic:

```go
if err := run(); err != nil {
    panic(err)
}
```

Then you can use done:

```go
done.Done(run())
```

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

---

## Features

- **Eliminates boilerplate**: Replace verbose `if err != nil` blocks with concise function calls
- **Type-safe validation**: Built with Go generics to provide compile-time type checking and IDE support
- **Chainable operations**: Link validation calls into clean, readable code flow
- **Multiple return types**: Handle functions returning values, pointers, slices, maps, and more
- **Rich assertions**: Specialized validators covering numbers, strings, booleans, and collections
- **Focus on logic**: Keep error handling lightweight so business logic stands out

---

## Installation

```bash
go get github.com/yyle88/done
```

---

## Usage

### Service Chain Validation

This example demonstrates three approaches: classic step-based validation, chained validation, and compact chained validation using done.

```go
package main

import (
	"fmt"

	"github.com/yyle88/done"
)

func main() {
	// Classic approach with explicit checks
	fmt.Println(WithClassicErrorHandling())

	// Chained approach with done.VCE
	fmt.Println(WithChainedErrorHandling())

	// Compact chained approach
	fmt.Println(WithCompactChainedHandling())
}

// WithClassicErrorHandling demonstrates classic validation with explicit checks.
// Pros: Distinct and simple to debug; Cons: Verbose code.
func WithClassicErrorHandling() string {
	service, err := NewService()
	if err != nil {
		panic(err) // Handle errors at each step
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
func WithChainedErrorHandling() string {
	service := done.VCE(NewService()).Nice()
	client := done.VCE(service.GetClient()).Nice()
	response := done.VCE(client.GetResponse()).Nice()
	return response.Message
}

// WithCompactChainedHandling shows the most compact form of chained validation.
// Pros: Concise code; Cons: Debugging is challenging.
func WithCompactChainedHandling() string {
	return done.VCE(done.VCE(done.VCE(
		NewService(),
	).Nice().GetClient(),
	).Nice().GetResponse(),
	).Nice().Message
}

// Service represents the main service in the chain.
type Service struct{}

// NewService creates a new Service instance.
func NewService() (*Service, error) {
	return &Service{}, nil
}

// GetClient returns a Client instance from this service.
func (s *Service) GetClient() (*Client, error) {
	return &Client{}, nil
}

// Client represents the intermediate client in the chain.
type Client struct{}

// GetResponse returns the Response containing the result message.
func (c *Client) GetResponse() (*Response, error) {
	return &Response{
		Message: "success", // Simulated success message
	}, nil
}

// Response represents the response containing the result.
type Response struct {
	Message string // Result message
}
```

⬆️ **Source:** [Source](internal/demos/demo3x/main.go)

### Chaining Operations

This example demonstrates two approaches to error handling: classic step-based validation and compact chained validation using done.

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/yyle88/done"
)

func main() {
	// Classic approach with explicit checks
	fmt.Println(WithClassicErrorHandling())

	// Compact chained approach
	fmt.Println(WithCompactChainedHandling())
}

// WithClassicErrorHandling demonstrates classic validation with explicit checks.
// Pros: Distinct and explicit; Cons: Verbose code.
func WithClassicErrorHandling() int64 {
	text, err := webFetch()
	if err != nil {
		panic(err) // Handle error at each step
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

// WithCompactChainedHandling uses done.VCE and done.VNE to provide concise validation.
// Pros: Compact code; Cons: Debugging can be challenging.
func WithCompactChainedHandling() int64 {
	// Chain methods to fetch, parse, and validate the value
	return done.VNE(
		parseNum(
			done.VCE(webFetch()).Nice(),
		),
	).Gt(0)
}

// webFetch simulates fetching a string value from a remote source.
func webFetch() (string, error) {
	return "100", nil // Simulated data fetch
}

// parseNum converts a string to an int64 value.
func parseNum(text string) (int64, error) {
	return strconv.ParseInt(text, 10, 64)
}
```

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

### Pointer Validation

This example demonstrates using done.P2 to validate and extract multiple pointer results.

```go
package main

import (
	"fmt"

	"github.com/yyle88/done"
)

func main() {
	// Classic approach: check results with explicit validation
	WithClassicErrorHandling()

	// Compact approach: use done.P2 to reduce boilerplate
	WithCompactErrorHandling()
}

// WithClassicErrorHandling demonstrates classic validation with explicit checks.
// Pros: Explicit and distinct; Cons: Verbose code.
func WithClassicErrorHandling() {
	account, config, err := fetchAccountAndConfig()
	if err != nil {
		panic(err) // Check each step
	}
	if account == nil {
		panic("account is nil") // Validate account exists
	}
	if config == nil {
		panic("config is nil") // Validate config exists
	}
	fmt.Println(account, config) // Print both account and config
}

// WithCompactErrorHandling uses done.P2 to streamline validation.
// Pros: Concise code; Cons: Validation is implicit.
func WithCompactErrorHandling() {
	account, config := done.P2(fetchAccountAndConfig()) // done.P2 handles checks
	fmt.Println(account, config)                        // Print both account and config
}

// Account represents an account in the system.
type Account struct {
	ID   int    // Account ID
	Name string // Account name
}

// Config represents configuration settings.
type Config struct {
	Timeout int    // Timeout in seconds
	Region  string // Service region
}

// fetchAccountAndConfig simulates fetching account and config data.
func fetchAccountAndConfig() (*Account, *Config, error) {
	account := &Account{ID: 1, Name: "Alice"}
	config := &Config{Timeout: 30, Region: "us-west"}
	return account, config, nil
}
```

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)

---

## Core Types

| Type                    | Description                                                                  |
|-------------------------|------------------------------------------------------------------------------|
| **`Ve[V any]`**         | Wraps any value with error. Provides `Done`, `Must`, `Soft` to handle errors |
| **`Vpe[V any]`**        | Wraps pointer with error. Validates non-nil with `Sure`, `Nice`, `Full` |
| **`Vce[V comparable]`** | Wraps comparable value with error. Compares values using `Same`, `Diff`, `Equals` |
| **`Vbe`**               | Wraps boolean with error. Asserts true/false with `TRUE`, `FALSE`, `OK`, `NO` |
| **`Vae[V any]`**        | Wraps slice with error. Checks emptiness and length with `Some`, `Have`, `Length` |
| **`Vme[K, V]`**         | Wraps map with error. Validates size and content with `Nice`, `Size`, `Len` |

---

## Core Functions

| Function   | Description                                    |
|------------|------------------------------------------------|
| **`Done`** | Panics with logging if error exists |
| **`Must`** | Ensures error is nil, returns value on success |
| **`Soft`** | Logs warning without panic, continues execution |
| **`Fata`** | Logs fatal error and terminates program |

---

## Function Groups

| Group                   | Functions                      | Description                                                            |
|-------------------------|--------------------------------|------------------------------------------------------------------------|
| **Errors**              | `Done`, `Must`, `Soft`         | Handle errors with panic/warning based on stage |
| **Non-Zero Validation** | `Sure`, `Nice`, `Some`         | Validate non-zero values and return them |
| **Non-Zero Assertion**  | `Good`, `Fine`, `Safe`         | Assert non-zero values without returning |
| **Zero Value Checking** | `Zero`, `None`, `Void`         | Validate values are zero/absent |
| **Value Comparisons**   | `Same`, `Diff`, `Is`, `Equals` | Compare values checking matches/differences |

---

## Type-Specific Validation

| Type      | Purpose | Methods |
|-----------|---------|---------|
| **`Vce`** | Comparable values | `Same`, `Diff`, `Is`, `Equals` - value comparison |
| **`Vse`** | String operations | `HasPrefix`, `HasSuffix`, `Contains` - substring checks |
| **`Vne`** | Numeric comparisons | `Gt`, `Lt`, `Gte`, `Lte` - range validation |
| **`Vbe`** | Boolean assertions | `TRUE`, `FALSE`, `YES`, `NO`, `OK` - true/false checks |
| **`Vae`** | Slice validation | `Sure`, `Some`, `Have`, `Length` - emptiness and size checks |
| **`Vme`** | Map validation | `Nice`, `Some`, `Size`, `Len` - map size and content checks |

---

## Examples

### Basic Error Handling

**Simple error check:**
```go
done.Done(run())
```

**Return value with error check:**
```go
result := done.V1(fetchData())
```

**Multiple return values:**
```go
v1, v2 := done.V2(getTwoValues())
```

### Pointer Validation

**Single pointer validation:**
```go
ptr := done.P1(getPointer())
```

**Multiple pointer validation:**
```go
ptr1, ptr2 := done.P2(getTwoPointers())
```

### Comparable Value Operations

**Check values are the same:**
```go
value := done.VCE(getValue()).Same(expected)
```

**Check values are different:**
```go
value := done.VCE(getValue()).Diff(unwanted)
```

### Numeric Comparisons

**Validate value exceeds threshold:**
```go
num := done.VNE(getNumber()).Gt(0)
```

**Less than validation:**
```go
num := done.VNE(getNumber()).Lt(100)
```

### Boolean Validation

**Ensure boolean is true:**
```go
done.VBE(checkCondition()).TRUE()
```

**Ensure boolean is false:**
```go
done.VBE(checkCondition()).FALSE()
```

### Slice Operations

**Ensure slice has elements:**
```go
items := done.VAE(getSlice()).Some()
```

**Check slice length:**
```go
items := done.VAE(getSlice()).Length(3)
```

### Map Operations

**Ensure map has content:**
```go
data := done.VME(getMap()).Nice()
```

**Check map size:**
```go
data := done.VME(getMap()).Size(5)
```

---

## Conclusion

The **Done** package brings a robust approach to error handling in Go. Through eliminating repetitive error checks, it lets you write clean, maintainable code. When building prototypes and production systems, Done helps you focus on what matters: the business logic.

**Test it out and let us know what you think!**

---

## Related Projects

Explore more error handling packages in this ecosystem:

### Advanced Packages

- **[must](https://github.com/yyle88/must)** - Must-style assertions with rich type support and detailed error context
- **[rese](https://github.com/yyle88/rese)** - Result extraction with panic, focused on safe value unwrapping

### Foundation Packages

- **[done](https://github.com/yyle88/done)** - Simple, focused error handling (this project)
- **[sure](https://github.com/yyle88/sure)** - Generates code that creates custom validation methods

Each package targets different use cases, from quick prototyping to production systems with comprehensive error handling.

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a mistake?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![starring](https://starchart.cc/yyle88/done.svg?variant=adaptive)](https://starchart.cc/yyle88/done)
