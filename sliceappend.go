package main

import "fmt"

func main() {

	x := []int{10, 20, 30, 40, 50}
	fmt.Println(x)

	x = append(x, 60, 70, 80)
	fmt.Println(x)

	y := []int{90, 100}

	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)

}
