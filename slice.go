package main

import "fmt"

func main() {

	x := []int{10, 20, 30, 40, 50}

	fmt.Println(x)

	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[2] = 25

	fmt.Println(x)

}
