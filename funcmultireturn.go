package main

import "fmt"

func main() {
	x, y := mouse("Ram", "Lakhan")
	fmt.Println(x)
	fmt.Println(y)
}
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, say "Hello"`)
	b := false
	return a, b
}
