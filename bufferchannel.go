package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 3)
	c <- 42
	fmt.Println(<-c)

	c <- 43
	c <- 44
	c <- 45
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
