package functions_test

import (
	"fmt"
)

const PI = 3.1416

func main3() {

	/*	numbers := []int{10, 12, 15, 17, 19, 22, 25, 26, 30, 34}
		for index, value := range numbers {
			fmt.Printf("Index is %d, value is %d\n", index, value)
		}*/

	/*sum, diff, product, div := calculator(18, 4)
	fmt.Println(sum, diff, product, div)*/

	/*	execute(5, 6, func(a int, b int) int {
		return a * b
	})*/

	/*incrementByThree := incrementer(3)
	incrementByThree(4)
	incrementByThree(9)

	incrementBySeven := incrementer(7)
	incrementBySeven(7)
	incrementBySeven(14)*/

	summation([]int{10, 15, 20, 25})
	summation2([]int{10, 15, 20, 25}...)
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

//function as return type

func incrementer(num int) func(int) {
	return func(x int) {
		fmt.Println(x + num)
	}
}

//variadic function

func summation(nums []int) {
	result := 0
	for _, num := range nums {
		result += num
	}
	fmt.Println(result)
}

func summation2(nums ...int) {
	result := 0
	for _, num := range nums {
		result += num
	}
	fmt.Println(result)
}

// annonymous function

// Practice-1
func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}

func PrintReports(intro, body, outro string) {
	printCostReport(func(intro string) int {
		return len(intro)
	}, intro)

	printCostReport(func(body string) int {
		return len(body)
	}, body)

	printCostReport(func(outro string) int {
		return len(outro)
	}, outro)
}

//Practice 2

func conversions(converter func(int) int, x, y, z int) (int, int, int) {
	convertedX := converter(x)
	convertedY := converter(y)
	convertedZ := converter(z)
	return convertedX, convertedY, convertedZ
}

func CallerFunction() {
	// using a named function
	//newX, newY, newZ := conversions(double, 1, 2, 3)
	// newX is 2, newY is 4, newZ is 6

	// using an anonymous function
	/* newX, newY, newZ = conversions(func(a int) int {
	    return a * a
	}, 1, 2, 3) */
	// newX is 1, newY is 4, newZ is 9
}

// Defer practice: Defer Scheduled a function-call to be executed later. commonly used for clean-up task i.e closing, file, closing DB connections, release resources, unlocking Mutex
func DeferPractice() {
	fmt.Println("Defer practicing ....")
	defer fmt.Println("Unlocking Mutex")
}

func sum(number1 int, number2 int) int {
	return number1 + number2
}

func calculateArea(radius float64) float64 {
	return PI * radius * radius
}

func calculatePerimeter(radius float64) float64 {
	return 2 * PI * radius
}

func calculateDiameter(radius float64) float64 {
	return 2 * radius
}

func CalculateCircleInDeclerativeStyle() float64 {
	var options int
	var result float64
	var radius float64
	fmt.Println("Enter the radius of the circle: ")
	fmt.Scanf("%f", &radius)
	fmt.Println("Choose Option \n 1. Area of circle \n 2. Perimeter of circle \n 3. Diameter of circle")
	fmt.Scanf("%d", &options)
	switch options {
	case 1:
		result = calculateArea(radius)
	case 2:
		result = calculatePerimeter(radius)
	case 3:
		result = calculateDiameter(radius)
	default:
		fmt.Println("Invalid Option entered.")
	}
	return result
}

func CalculateCircle() {
	// fmt.Println(" result is: ", CalculateCircleInDeclerativeStyle())
	var options int
	var radius float64
	fmt.Println("Enter the radius of the circle: ")
	fmt.Scanf("%f", &radius)
	fmt.Println("Choose Option \n 1. Area of circle \n 2. Perimeter of circle \n 3. Diameter of circle")
	fmt.Scanf("%d", &options)
	fun := getFunction(options)
	CalculateCircleInFunctionalStyle(radius, fun)
}

func getFunction(query int) func(float64) float64 {
	functions := map[int]func(float64) float64{
		1: calculateArea,
		2: calculatePerimeter,
		3: calculateDiameter,
	}
	return functions[query]
}

func CalculateCircleInFunctionalStyle(radius float64, f func(float64) float64) {
	result := f(radius)
	fmt.Println("Result: ", result)
	fmt.Println("Thank you!!")
}

// 