package interface_test

import (
	"fmt"
	"strconv"
	"strings"
)

/**
Struct topic covered
	1. creating struct
	2. accessing field
	3. modifying field
	4. method on struct
*/

type IntSlice []int

func (intSlice IntSlice) String() string {
	var strs []string
	for _, v := range intSlice {
		strs = append(strs, strconv.Itoa(v))
	}
	return "[" + strings.Join(strs, ";") + "]"
}

func IntSliceDemo() {
	var v IntSlice = []int{1, 2, 3}
	for i, x := range v {
		fmt.Printf("%d %d \n", i, x)
	}
	fmt.Printf("%T %[1]v\n", v)
}

func main4() {
	/*
		var student = Student{"Alauddin Tuhin", 27, "eu.tuhin@gmail.com", "16"}
		fmt.Println(student)

		//2. accessing field
		fmt.Printf("Name: %s Age: %d Email: %s\n", student.name, student.age, student.email)

		//3. modifying field
		student.name = "Mr. Tuhin Md Alauddin"
		fmt.Printf("Modified Name: %s\n", student.name)

		student2 := Student{email: "alutuhin@gmail.com", age: 28}
		student2.printData()

		employee := struct {
			firstName string
			lastName  string
			age       int32
			salary    int32
		}{"Alauddin", "Tuhin", 28, 50000}

		fmt.Println("First Employee: ", employee)*/

	fmt.Println(divide(4, 0))
}

// 1. creating struct
type Student struct {
	name  string
	age   int
	email string
	roll  string
}

// 4. method on struct
func (student Student) printData() {
	fmt.Printf("Name: %s Age: %d Email: %s\n", student.name, student.age, student.email)
}

func divide(a, b int) int {

	return a / b
}
