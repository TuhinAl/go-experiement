package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Alauddin Tuhin ") // This is a comment

	// ===================== GO variables ==================
	var x float64 = 20.0

	fmt.Println(" The value of x-is: ", x)
	fmt.Printf("x is of type %T\n", x)
	var num int
	var amount float32
	var name string
	var isValid bool
	/*num = 20
	amount = 399.99
	isValid = false
	name = "Tuhin"*/
	fmt.Println(num, amount, isValid, name)

	y := "Alauddin Tuhin" // Shorthand variable declare and initialize
	fmt.Printf("y type of %T\n", y)

	value1, value2 := 55, "hi"
	fmt.Println(value1, value2)

	value3, value2 := true, "hello"

	fmt.Println(value3, value2)

	if value4, value2 := 66, "hey"; value4 > value1 {
		fmt.Println("value4 and value1: = ", value4, value2)
	}
	fmt.Println(value2)
	//fmt.Println(value4)

}
