package main

import (
	"fmt"
)

func main() {
	/*fmt.Println("Hello World!")
	fmt.Println("Alauddin Tuhin ") // This is a comment

	// ===================== GO variables ==================
	var x float64 = 20.0

	fmt.Println(" The value of x-is: ", x)
	fmt.Printf("x is of type %T\n", x)
	var num int
	var amount float32
	var name string
	var isValid bool
	num = 20
	amount = 399.99
	isValid = false
	name = "Tuhin"
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
	fmt.Println(value4)*/

	fmt.Println("======================== Array=============================")
	/*var array = [10]string{"Real Madrid", "Arsenal", "Manchester United"}
	var teams = [5]string{"Real Madrid", "Arsenal", "Manchester United", "Bayern Manchen", "PSG"}
	var epl = [5]string{"Arsenal", "Manchester United", "Chelsea", " Liverpool", "Tottenham"}
	var seria = epl
	var seria2 = &epl
	fmt.Println(seria == seria2) Invalid type because mismatch type
	var laLiga = [5]string{"Real Madrid","Barcelona", "Chelsea", "Sevilla", "Atletico Madrid""}
	var championes = [10]string{"Real Madrid", "Arsenal", "Manchester United", "Bayern Manchen", "PSG", "Barcelona", "Chelsea", " Milan", "Napoli"}
	var championes2 = [10]string{"Real Madrid", "Arsenal", "Manchester United", "Bayern Manchen", "PSG", "Barcelona", "Chelsea", " Milan", "Napoli"}
	var array2 = [...]string{"Real Madrid", "Arsenal", "Manchester United"}
	var array3 = [...]string // not possible to declare like this

	fmt.Println(array)
	fmt.Println(reflect.ValueOf(array).Kind())
	fmt.Println("Array length: ", len(array))
	fmt.Println("Array length2: ", len(array2))
	//fmt.Println("Array length2: ", len(array3))

	// Initializing Array with value at specific index
	village := [10]int{1: 163, 4: 188, 8: 255}
	x := [5]int{1: 10, 3: 30}
	fmt.Println(len(x))
	fmt.Println(village)*/

	/*for i := 0; i < len(teams)-1; i++ {
		fmt.Printf("Rank-> %d and Team is %s\n", i, teams[i])
	}

	for index, element := range teams {
		fmt.Println(index, "==>", element)
	}

	for _, values := range teams {
		fmt.Println(values)
	}
	i := 0
	for range championes {
		fmt.Println(teams[i])
		i++
	}

	s1 := "Dhaka"
	s2 := "Dhaka"

	println("champions address: ", &championes)
	println("champions2 address: ", &championes2)
	println("are the equal: ", &championes2 == &championes)
	println("are the equal: ", championes2 == championes)
	println("are the equal: ", &s1 == &s2)
	println("are the equal: ", s1 == s2)*/

	/*var name1 string
	name2 := "tuhin"
	var_1, var_2 := 1, "hi" //declare var_1 int and var_2 string
	fmt.Println(var_1, var_2)
	ar_3, var_2 := 2, "hello"
	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(ar_3)
	fmt.Println(var_2)
	fmt.Printf("type of name1 is: %T\n", name1)*/

	var1, var2, var3 := 3, 4, "Dhaka"
	fmt.Println("var3: ", var3)

	if var1 < var2 {
		fmt.Print("inside if condition.")
		var4 := 100
		var3 := "Khulna"
		fmt.Println("var3(inside scope): ", var3)
		fmt.Println("var3: ", var4)
	}
	fmt.Println("var3: ", var3)
	//fmt.Println("var4: ", var4)

}
