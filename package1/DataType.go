package main

import (
	"fmt"
)

func main() {
	fmt.Println("================================= data type =============================")

	/*	there are two types of data type
			1. Basic Type
				(i) Numeric
		 		(ii) String
				(iii) Boolean

			2. Composite Type
	*/

	str := "Welcome to the Go Programming language"
	id := 350
	result := 3.63
	//var res int8;
	boolean := false
	//char := 'T'
	//println("===========String Datatype ======")
	//fmt.Printf("%s\n", str)
	//fmt.Printf("%q\n", str)
	//fmt.Printf("%x\n", str)
	//fmt.Printf("%X\n", str)

	//println("===========int Datatype ======")
	//fmt.Printf("%d\n", id)
	//fmt.Printf("%b\n", id)
	//fmt.Printf("%o\n", id)
	//fmt.Printf("%x\n", id)
	//fmt.Printf("%X\n", id)

	//println("===========float Datatype ======")
	//fmt.Printf("%f\n", result)
	//fmt.Printf("%e\n", result)
	//fmt.Printf("%E\n", result)
	//fmt.Printf("%g\n", result)
	fmt.Printf("%t\n", boolean)

	println("===========printing pointer address ======")
	fmt.Printf("%p\n", &str)
	fmt.Printf("%p\n", &id)
	fmt.Printf("%p\n", &result)
	fmt.Printf("%p\n", &boolean)

	// Signed and unsigned integer

	//Signed means positive and negative numbers { -5 to 5}
	// unsigned means only positive numbers { 0 to 5}

}
