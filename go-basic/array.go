package main

import "fmt"

func main() {
	//arrayDemo()
	sliceDemo()
}

func arrayDemo() {

	var division [3]string
	division = [3]string{"Dhaka", "Khulna", "Sylhet"}
	fmt.Println(division)

	fruits := [...]string{"Mango", "Banana", "Pineapple", "Guava", "Apple"}
	fmt.Printf("length of fruits: %d\n ", len(fruits))

	/*for i := 0; i < len(fruits); i++ {
		fmt.Println("fruit: ", fruits[i])
	}*/

	for i, value := range fruits {
		fmt.Printf("fruits[%d]: %s\n", i, value)
	}

	for _, value := range fruits {
		fmt.Printf("fruits: %s\n", value)
	}
}

func sliceDemo() {
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5, 6)
	fmt.Println(slice)

	slice2 := make([]int, len(slice))
	copy(slice2, slice)
	fmt.Println(slice2)
	slice2 = append(slice2, 16, 17, 18)
	fmt.Println(slice)
	fmt.Println(slice2)
}
