package datatypeoperator

import (
	"fmt"
	"strings"
)

/*
& -> Address of operator
* -> Dereference operator
A pointer value is the address of a variable. A pointer is thus the location at which a value is stored. Not every value has an address, but every variable does.

&x => address of x
*int => pointer of int
p := &x   => p points to x, or p contains the address of x.’

When you declare var p *int, it means p can hold the memory address of an integer value.




Example:
 	x := 42               // x is an integer variable
    p := &x               // p is a pointer to x

    1. fmt.Println("x:", x)  // Value of x
    2. fmt.Println("p:", p)  // Address of x
    3. fmt.Println("*p:", *p) // Value at the address p points to (dereferencing)

    *p = 100              // Change the value at the address p points to
    4. fmt.Println("x after change:", x) // x is updated

Output:
	1. x: 42
	2. p: 0xc00001a088
	3. *p: 42
	4. x after change: 100

*/

type Student struct {
	Id    int
	Name  string
	age   int
	email string
}

func PointerPractice() {
	tuhin := Student{Id: 1, Name: "Alauddin Tuhin", age: 28}
	modifyStudentStruct(&tuhin, "gmail")
	fmt.Println(tuhin.Id)
	fmt.Println(tuhin.Name)
	fmt.Println(tuhin.age)
	fmt.Println(tuhin.email)
	value := 13
	fmt.Println(passByValue(&value))
	fmt.Println(returnAddresss() == returnAddresss()) // false
	value = 17
	var pointerValue *int
	pointerValue = &value

	fmt.Println(pointerIncrement(pointerValue))
	fmt.Println(value)

}

func modifyStudentStruct(student *Student, domain string) {
	student.Id += 5
	student.age += 7
	student.Name = "Mr./Mrs " + strings.ToUpper(student.Name)
	student.email = strings.ToLower(strings.ReplaceAll(strings.TrimPrefix(student.Name, "Mr./Mrs "), " ", "")) + "@" + domain + ".com"
}

func passByValue(data *int) int {
	a := 11
	*data = a
	return *data
}

func returnAddresss() *int {
	value := 15
	return &value
}

func pointerIncrement(point *int) int {
	fmt.Println("*point: ", *point)
	*point++
	return *point
}
