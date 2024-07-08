package main

import "fmt"

func main() {

	b := 125
	var a *int = &b

	fmt.Printf("Type of a is: %T\n", a)
	fmt.Println("Address of b is: ", a)

	c := 125
	var d *int

	if d == nil {
		fmt.Println("d is: ", d)
		d = &c
		fmt.Println("d after initialization is", d)
	}

}
