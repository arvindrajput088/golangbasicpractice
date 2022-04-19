package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	age       int
}

func main() {
	p1 := person{
		firstname: "Arvind",
		lastname:  "Rajput",
		age:       31,
	}
	p2 := person{
		firstname: "Deep",
		lastname:  "Rajput",
		age:       29,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.firstname, p1.lastname, p1.age)
	fmt.Println(p2.firstname, p2.lastname, p2.age)

}
