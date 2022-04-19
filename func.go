package main

import "fmt"

func main() {
	foo()
	bar("Arvind")
	s1 := boo("Rajput")
	fmt.Println(s1)
}

func foo() {

	fmt.Println("Hello")
}

func bar(s string) {

	fmt.Println("Hii,", s)
}

func boo(st string) string {
	return fmt.Sprint(st, " ", "Hello from woo  ")
}
