package main

import (
	"fmt"
)

func main() {

	//slice();
	//maps()
	//pointersDemo()
	pointerWithArray()
}

func slice() {
	//var array [5]int
	array := [...]int{10, 20, 30, 40, 50}
	fmt.Println(array)
	fmt.Println(len(array))
	fmt.Printf("%T\n", array)

	//Slice creating 3 ways

	mySlice := []int{15, 20, 25, 30, 35, 40}
	sliceFromArray := array[2:5]
	sliceUsingMake := make([]int, 15)
	fmt.Println(mySlice)
	fmt.Println(sliceFromArray)
	fmt.Println(sliceUsingMake)
	fmt.Println(len(mySlice))
	//fmt.Println(cap(mySlice))
	// Slice Operation
	//append
	mySlice = append(mySlice, 20, 21, 22, 23)
	fmt.Println(len(mySlice))

	//iterate
	for index, value := range mySlice {
		fmt.Printf("Indexis %d, and value is %d\n", index, value)
	}

	//copy (Shallow copy)
	/*newMySlice := mySlice
	fmt.Println(mySlice)
	fmt.Println(newMySlice)
	newMySlice[1] = 111
	fmt.Println(mySlice)
	fmt.Println(newMySlice)*/
	//copy (Deep copy)
	newMySlice := make([]int, len(mySlice))
	copy(newMySlice, mySlice)
	fmt.Println(newMySlice)
}

func maps() {
	var maps = map[string]int{
		"Dhaka":      1,
		"Chittagong": 2,
		"Khulna":     3,
		"Rajshahi":   4,
		"Sylhet":     5,
		"Rangpur":    6,
	}
	/*fmt.Println(maps)
	maps["Mymensingh"] = 7
	fmt.Println(maps)
	maps["Khulna"] = 2
	fmt.Println(maps)*/

	//iterating
	for key, value := range maps {
		fmt.Printf("Key is: %s and value is: %d\n", key, value)
	}

	var person Person = Person{"Alauddin", "Tuhin", 28, "alauddintuhin1411@gmail.com"}
	fmt.Println(person)
}

type Person struct {
	firstName string
	lastName  string
	age       int
	email     string
}

func pointersDemo() {
	var num int
	num = 10
	fmt.Println(num, &num)

	var numPointer *int
	numPointer = &num
	fmt.Println(*numPointer, &numPointer)
	*numPointer = 15
	fmt.Println(*numPointer, num)

	// new and make keyword in golang
	/**
	new() => allocates memory for a new zeroed value of a specified type and returns a pointer to it.It is primarily
	used for initializing and obtaining a pointer to a newly allocated zeroed value of a given type, usually for data
	types like structs. new() basically work with int,float, boolean type data and make() work with slice, map, channel.

	make() => used for initializing slices, maps, and channels – data structures that require runtime initialization.
	Unlike new(), make() returns an initialized (non-zeroed) value of a specified type.

	Pointer vs. Value:
	Keep in mind that new() returns a pointer, while make() returns a non-zeroed value. Choose the appropriate method
	based on whether you need a pointer or an initialized value.
	*/

	// pointers with array

}

func pointerWithArray() {
	array := []int{10, 11, 12, 13, 14}
	fmt.Println(&(array[0]))
	modifyArray(&(array))
	fmt.Println(array)
}
func modifyArray(a *[]int) {
	fmt.Println((*a)[0], &(*a)[0])
	//a[0] = 22
	fmt.Println(*a)
}
