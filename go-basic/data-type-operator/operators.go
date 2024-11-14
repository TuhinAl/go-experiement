package datatypeoperator

import (
	"fmt"
)

func main6() {

	// if else
	/*num := -11
	if num > 0 {
		fmt.Println("Is Positive")
	} else if num == 0 {
		fmt.Println("Is Zero")
	} else {
		fmt.Println("Is Negative")
	}*/

	//switch practice
	/*var day string
	day = "sunday"
	switch day {
	case "saturday", "sunday":
		println("Weekend.")
	case "friday", "monday", "tuesday", "wednesday", "thursday":
		println("Workday!!")
	default:
		println("Invalid Day.")
	}*/

	//for loop (normal,infinite, range)

	/*	for i := 0; i < 5; i++ {
			println("Bangladesh!!")
		}

		for {
			println("Bangladesh!!");
		}*/

	numbers := []int{10, 12, 15, 17, 19, 22, 25, 26, 30, 34}
	for index, value := range numbers {
		fmt.Printf("Index is %d, value is %d\n", index, value)
	}
}
