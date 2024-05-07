# Generic Fixture Library for Go

A flexible and robust library for creating test fixtures in Go. This library simplifies the process of constructing objects for testing by allowing you to easily set up test data through a fluent API.

## Features

- **Type-Safe**: Utilizes Go generics to ensure that your fixtures are type-safe.
- **Flexible**: Modify properties of your objects easily at runtime.
- **Error Handling**: Provides detailed error reports if a field is invalid or a type mismatch occurs.

## Getting Started

To use this library, simply import it into your Go project.

```bash
go get github.com/pers0na2dev/fixture@latest
```

Example Usage

Here is a quick example to show you how to use the library:

```go
package main

import (
    "fmt"

    "github.com/pers0na2dev/fixture"
)

type User struct {
    Name string
    Age  int
}

func main() {
    userFixture := fixture.NewFixture[User](
        fixture.With{"Name", "John Doe"},
        fixture.With{"Age", 30},
    )

    user, err := userFixture.Build()
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("User:", user)
    }
}
```
