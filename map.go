package main

import "fmt"

func main() {
	x := map[string]int{
		"ram":    1000,
		"lakhan": 5000,
		"deep":   5000,
	}
	fmt.Println(x)
	fmt.Println(x["ram"])
	fmt.Println(x["arv"])
	v, ok := x["ram"], false
	fmt.Println(v, ok)
	fmt.Println(ok)
	x["pri"] = 6000
	fmt.Println(x)
	delete(x, "ram")
	fmt.Println(x)

}
