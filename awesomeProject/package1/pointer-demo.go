package main

import (
	//"structs/computer"
	"fmt"
	//"structs/computer"
)

func main() {

	/*b := 125
	var a *int = &b

	fmt.Printf("Type of a is: %T\n", a)
	fmt.Println("Address of b is: ", a)

	c := 125
	var d *int

	if d == nil {
		fmt.Println("d is: ", d)
		d = &c
		fmt.Println("d after initialization is", d)
	}*/

	e := 60
	fmt.Printf("value of a before function call is: %T\n", e)

	f := &e
	valueChange(f)
	fmt.Printf("value of a after function call is: %T\n", f)

	/*		spec := computer.Spec{
				Maker: "Apple",
				Price: 50000,
			}

			fmt.Println("Maker: ", spec.)*/

}

func valueChange(val *int) {
	*val = 55
}
