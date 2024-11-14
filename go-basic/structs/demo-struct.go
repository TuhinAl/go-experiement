package structs

import "fmt"

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
