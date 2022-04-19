package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (a person) speak() {
	fmt.Println("I am", a.first, a.last, "and i am", a.age, "year old")
}
func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p1.speak()
}
