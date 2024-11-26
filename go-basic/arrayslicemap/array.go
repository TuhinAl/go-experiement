package arrayslicemap

import "fmt"

/* func main() {
	//arrayDemo()
	sliceDemo()
}
*/

// iota is special constant geenrator that automatically assigns increasing integer values to constants, starting from 0
type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	BDT
)

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

func arrayBasic() {
	footballTeams := [...]string{"Real Madrid", "Manchester City", "Bayern Munchen", "Barcelona", "Liverpool", "PSG"}
	fmt.Printf("%T\n", footballTeams)
}
func currencySymbol() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", BDT: "৳"}
	//above array look like this=> "symbol := [4]string{"$", "€", "£", "৳"}""
	fmt.Println(BDT, symbol[BDT])
	testArray := [...]string{0: "Dhaka", 1: "Khulna", 3: "Rajshahi"}
	fmt.Println(testArray[1])
	initialize100SizedArray := [...]int{99: -1}
	fmt.Println(initialize100SizedArray[99])
}

func arrayComparableCheck() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	// d := [3]int{1,2}
	fmt.Println(a == b, a == c, b == c) //true false false
	// fmt.Println(a==d) // compile error
}

func ArrayPractice() {
	// arrayBasic()
	// currencySymbol()
	arrayComparableCheck()
}
