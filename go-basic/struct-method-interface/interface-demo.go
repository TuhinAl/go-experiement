package interface_test

import (
	"fmt"

	"github.com/google/uuid"
)

func main9() {
	/*	name := MyString("alauddin tuhin")
		var vowelsFinder VowelsFinder
		vowelsFinder = name // possible since MyString implements VowelsFinder
		fmt.Printf("Vowels are %c", vowelsFinder.FindVowels())*/

	employees := sliceOfSalaryCalculator()
	totalExpense(employees)
	//fmt.Printf("Total Expenses are %d", total)
}

type VowelsFinder interface {
	FindVowels() []rune // rune is alias of int 32
}

type MyString string

//MyString implements VowelsFinder

func (myString MyString) FindVowels() []rune {
	var vowels []rune

	for _, vowel := range myString {
		if vowel == 'a' || vowel == 'e' || vowel == 'i' || vowel == 'o' || vowel == 'u' {
			vowels = append(vowels, vowel)

		}
	}
	return vowels
}

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	employeeId  int
	basicSalary int
	pf          int
}

type Contract struct {
	employeeId  int
	basicSalary int
}

type Freelancer struct {
	empId       int
	ratePerHour int
	totalHours  int
}

func (permanent Permanent) CalculateSalary() int {
	//grantPF := permanent.basicSalary * .07;
	return permanent.basicSalary + permanent.pf
}

func (contract Contract) CalculateSalary() int {
	return contract.basicSalary
}

func (freelancer Freelancer) CalculateSalary() int {
	return freelancer.ratePerHour * freelancer.totalHours * 22
}

func totalExpense(salaryCalculator []SalaryCalculator) {
	totalExpense := 0
	for _, salary := range salaryCalculator {
		totalExpense += salary.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month ৳%d", totalExpense)
}

func sliceOfSalaryCalculator() []SalaryCalculator {
	selim := Permanent{100, 50000, 3500}
	josim := Permanent{101, 45000, 2700}
	bapparaj := Contract{102, 70000}
	tanvir := Freelancer{103, 340, 8}
	freelancer2 := Freelancer{104, 220, 9}
	return []SalaryCalculator{selim, josim, bapparaj, tanvir, freelancer2}
}

type Product struct {
	Id       uuid.UUID
	Name     string
	Price    float64
	Quantity int
}
type ProductDto struct {
	Name     string
	Price    float64
	Quantity int
}

type ProductService interface {
	SaveProduct(p ProductDto) ProductDto
	EditProduct(p ProductDto) ProductDto
	DeleteProduct() Product
	SearchProduct() Product
}
type ProductServiceImpl struct{}

func (s ProductServiceImpl) SaveProduct(p ProductDto) ProductDto {
	product := Product{
		Id:       uuid.New(),
		Name:     p.Name,
		Price:    p.Price,
		Quantity: p.Quantity,
	}
	return ProductDto{Name: product.Name}
}

// func (p *ProductServiceImpl) EditProduct(dto ProductDto) ProductDto { // pointer receiver
// 	p.Price = dto.Price
// 	return p
// }

func ImplementationTest() {
	banana := ProductDto{
		Name:     "Banana",
		Price:    10.50,
		Quantity: 8,
	}
	apple := ProductDto{
		Name:     "Apple",
		Price:    75.20,
		Quantity: 8,
	}
	orange := ProductDto{
		Name:     "Orange",
		Price:    60.00,
		Quantity: 8,
	}

	var productService ProductServiceImpl
	bananaDto := productService.SaveProduct(banana)
	appleDto := productService.SaveProduct(apple)
	orangeDto := productService.SaveProduct(orange)
	fmt.Println(bananaDto)
	fmt.Println(appleDto)
	fmt.Println(orangeDto)

}


func InterfaceExplore()  {
	// var shape Shape
	rectangleShape := Rectangle{width: 15.00, length: 10.75}
	triangleShape :=  Triangle{width: 8.30, height: 7.50}
	circleShape := Circle{Radius: 7.75}
	fmt.Println(rectangleShape.Area())
	fmt.Println(triangleShape.Area())
	fmt.Println(circleShape.Area())
	fmt.Println(circleShape.Perimeter())
	fmt.Println(circleShape.Diameter())
}

type Shape interface{
	Area() float64
	Perimeter()float64
	Diameter()float64
}

type Circle struct{
	Radius float64
}
type Rectangle struct{
	width, length float64
}
type Triangle struct{
	width, height float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.1416
}

func (c Triangle) Area() float64 {
	return c.width * c.height * 0.5
}

func (r Rectangle) Area() float64 {
	return r.width * r.length
}

func (c Circle) Diameter() float64 {
	return 2 * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * 3.1416
}