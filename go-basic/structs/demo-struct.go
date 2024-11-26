package structs

import (
	"fmt"
	"time"
)

type Car struct {
	string
	int
	//promoted fields
	Person
}

// Nested Structs
type Person struct {
	name    string
	email   string
	age     int32
	address Address
}

type Address struct {
	city     string
	district string
	country  string
}

// Spec Struct exported
type Spec struct {
	Maker string // exported field
	Price int    // exported field
	model string
}
type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerId     int
}

func structMethod() {
	person := Person{
		name:  "Jasim Uddin",
		email: "jasim@email.com",
		age:   25,
		address: Address{
			city:     "New York",
			district: "New York",
			country:  "USA",
		},
	}

	fmt.Println("Name:", person.name)
	fmt.Println("Email:", person.email)
	fmt.Println("Age:", person.age)
	fmt.Println("city:", person.address.city)
	fmt.Println("district:", person.address.district)
	fmt.Println("country:", person.address.country)
}

func EmployeeManipulation() {
	atikurRahman := Employee{
		ID:        1,
		Name:      "Atiqur Rahman",
		Address:   "Rajshahi, Bangladesh",
		DoB:       time.Now(),
		Position:  "Juniro Software Enginner",
		Salary:    22800,
		ManagerId: 1000,
	}
	fmt.Println("Before Salary Increment: Atiq Rahman is: ", atikurRahman.Salary)
	promotionAndSalaryIncrement(&atikurRahman, "Senior Software Engineer", 5000).Salary = 0
	fmt.Println("Updated Atiq: Atiq Rahman is: ", atikurRahman.Salary)
}

func BasicStruct() {
	// comparing struct

	// Struct Embedding
	var wheel1, wheel2 Wheel
	// wheel.X = 10
	// wheel.Y = 8
	wheel1 = Wheel{Circle: Circle{
		Point: Point{X: 11, Y: 12}, Radius: 5,
	}, Spoke: 10}
	wheel2 = Wheel{Circle: Circle{
		Point: Point{X: 11, Y: 12}, Radius: 5,
	}, Spoke: 10}
	fmt.Println(wheel1 == wheel2)
}

// todo: will validate and Error check
func promotionAndSalaryIncrement(employee *Employee, designation string, incrementAmount int) *Employee {
	employee.Salary += incrementAmount
	currentPosition := &employee.Position
	*currentPosition = designation
	return employee
}

// Binary tree to implement an insertion sort:
type Tree struct {
	data        int
	left, right *Tree
}

func Sort(values []int) {
	var root *Tree
	for _, v := range values {
		root = add(root, v)
	}

	appendValues(values[:0], root)
}

func appendValues(values []int, root *Tree) []int {
	if root != nil {
		values = appendValues(values, root.left)
		values = append(values, root.data)
		values = appendValues(values, root.right)
	}
	return values
}

func add(root *Tree, value int) *Tree {
	if root == nil {
		root = new(Tree)
		root.data = value
		return root
	}

	if value < root.data {
		root.left = add(root.left, value)
	} else {
		root.right = add(root.right, value)
	}
	return root
}

type Point struct {
	X int
	Y int
}
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spoke int
}
