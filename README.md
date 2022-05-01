# 🚀 GoDI: Generic based DI in Go

## Installation

`DI`:

* `go get github.com/MaximZayats/godi/`
* ```go
  import "github.com/MaximZayats/godi/di"
  ```

`CMD`:

* ```shell
  go get github.com/MaximZayats/godi/cmd/godi
  ```
* ```shell
  go run github.com/MaximZayats/godi/cmd/godi init ./distorage
  ```
* ```shell
  go run github.com/MaximZayats/godi/cmd/godi --help
  ```

## Example

* **Simple**: Getting from container
    ```go
    package main
    
    import (
        "fmt"
        "github.com/MaximZayats/godi/di"
    )
    
    func main() {
        di.AddSingletonByFactory[float32](func(c *di.Container) float32 {
            return 11.22
        })
    
        di.AddInstance[int](123)
    
        di.AddScopedByFactory[string](func(c *di.Container) string {
            return "aabbcc"
        })
    
        fmt.Println(di.Get[int]())     // 123
        fmt.Println(di.Get[string]())  // "aabbcc"
        fmt.Println(di.Get[float32]()) // 11.22
        fmt.Println(di.Get[float32]()) // 11.22
    }
    ```

* **Injection**: Pass arguments to function (*Simple code generation is required*)

  **Full code**: [godi-fiber-example](https://github.com/MaximZayats/godi-fiber-example)

  **Snippet**:
  ```go
  type H = func(*fiber.Ctx) error

  // `stringFromDI` will be injected into the handler
  func handler(c *fiber.Ctx, stringFromDI string) error {
      return c.SendString("Hello from di: " + stringFromDI)
  }
  
  func main() {
      di.AddInstance[string]("I'm string from DI!!!", c)
      ...
      app.Get("/", injection.Inject[H](handler))
  }
  ```

[Other examples](/examples)

## Usage

### TODO

## Benchmarks

[Code](/benchmark/local_container_test.go)

```text
goos: windows
goarch: amd64
pkg: github.com/MaximZayats/godi/benchmark
cpu: AMD Ryzen 5 1600 Six-Core Processor
BenchmarkGetFromFactorySingleton
BenchmarkGetFromFactorySingleton-12     500488393                2.443 ns/op
BenchmarkGetInstance
BenchmarkGetInstance-12                 495795447                2.403 ns/op
BenchmarkGetFromFactory
BenchmarkGetFromFactory-12              361722957                3.273 ns/op
PASS
```
