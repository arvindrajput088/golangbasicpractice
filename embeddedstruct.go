package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	age       int
}
type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			firstname: "Arvind",
			lastname:  "Rajput",
			age:       31,
		},
		ltk: true,
	}
	p2 := person{
		firstname: "Deep",
		lastname:  "Rajput",
		age:       29,
	}
	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Println(sa1.firstname, sa1.lastname, sa1.age, sa1.ltk)
	fmt.Println(p2.firstname, p2.lastname, p2.age)

}
