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
	updatedAtiq := promotionAndSalaryIncrement(&atikurRahman, "Senior Software Engineer", 5000).Salary = 0
	fmt.Println("Updated Atiq: Atiq Rahman is: ", atikurRahman.Salary)
}

// todo: will validate and Error check
func promotionAndSalaryIncrement(employee *Employee, designation string, incrementAmount int) * Employee {
	employee.Salary += incrementAmount
	currentPosition := &employee.Position
	*currentPosition = designation
	return *employee
}
