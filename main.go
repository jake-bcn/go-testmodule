// main.go
package main

import (
	"fmt"

	test "github.com/jake-bcn/go-testmodule/test"
)

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	fmt.Println(Greet("World"))
	fmt.Println(test.Greet("World"))
}
