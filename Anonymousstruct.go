package main

import "fmt"

func main() {
	p1 := struct {
		firstname string
		lastname  string
		age       int
	}{
		firstname: "Arvind",
		lastname:  "Rajput",
		age:       31,
	}

	fmt.Println(p1)
}
