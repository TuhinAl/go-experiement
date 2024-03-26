package main

import "fmt"

func main() {

	/*	numbers := []int{10, 12, 15, 17, 19, 22, 25, 26, 30, 34}
		for index, value := range numbers {
			fmt.Printf("Index is %d, value is %d\n", index, value)
		}*/

	/*sum, diff, product, div := calculator(18, 4)
	fmt.Println(sum, diff, product, div)*/

	execute(5, 6, func(a int, b int) int {
		return a * b
	})

}

func addition(number1 int, number2 int) int {
	return number1 + number2
}

func calculator(input1 int, input2 int) (sum int, diff int, product int, div float64) {
	sum = input1 + input2
	diff = input1 - input2
	product = input1 * input2
	div = float64(input1 / input2)
	return
}

func execute(x int, y int, operation func(x int, y int) int) {
	result := x + y + operation(x, y)
	fmt.Println(result)
}
