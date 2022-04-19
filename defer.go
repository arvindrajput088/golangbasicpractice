package main

import (
	"fmt"
)

func main() {
	defer foo()
	bar()
	abc()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
func abc() {
	fmt.Println("abc")
}
