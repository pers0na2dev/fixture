package main

import (
	"fmt"

	"github.com/pers0na2dev/fixture"
)

type Nested struct {
	Another int
}

type MyStruct struct {
	name int // works with unexported fields
	Some int // works with exported fields
	Nested Nested // works with nested structs
}

func main() {
	f := fixture.NewFixture[MyStruct](
		fixture.With{Name: "name", Value: 10},
		fixture.With{Name: "Some", Value: 20},
		fixture.With{Name: "Nested", Value: Nested{Another: 30}},
	)
	/* or
	// f.With("name", 10)
	*/

	v, err := f.Build()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(v)
}