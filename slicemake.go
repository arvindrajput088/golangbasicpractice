package main

import "fmt"

func main() {
	x := make([]int, 5, 10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[2] = 3
	x[3] = 4
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 6, 7)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
