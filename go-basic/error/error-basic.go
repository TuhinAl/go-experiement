package error__test

import (
	"errors"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)
import "C"

/**
//errors are value in go
topic covered:
Understanding Error interface
Importance of error checking
Custom error, error Wrapping
Defer Panic and recover
*/

func main() {
	//testDefer()
	//openFile()
	//convertToUndelyingTypeError()
	dnsErrorCheck()
}

//=========================================== Defer===========================

func totalTime(startTime time.Time) {
	fmt.Printf("Total time taken: %f seconds", time.Since(startTime).Seconds())
}

func testDefer() {
	/* startTime := time.Now()
	defer totalTime(startTime)
	time.Sleep(2 * time.Second)
	fmt.Println("Sleep complete") */

	//defer argument evaluation
	/* a := 5
	defer displayValue(a)
	a = 10
	fmt.Println("value of a before deferred function call", a) */

	//defer in method
	/* person := Person {
		firstName: "John",
		lastName: "Smith",
	}

	defer person.fullName()
	fmt.Printf("Welcome ")	 */

	// pactical defer in real life

	var waitGroup sync.WaitGroup

	rectangular1 := Rectangular{10, 12}
	rectangular2 := Rectangular{-17, 3}
	rectangular3 := Rectangular{8, 12}
	rectangular4 := Rectangular{11, 11}
	rectangles := []Rectangular{rectangular1, rectangular2, rectangular3, rectangular4}
	for _, v := range rectangles {
		waitGroup.Add(1)
		go v.area(&waitGroup)
		//fmt.Println("area: %d\n", result)
	}
	waitGroup.Wait()
	fmt.Println("All go routines finished executing")
}

type Person struct {
	firstName string
	lastName  string
}

func (person Person) fullName() {
	fmt.Printf(" : %s %s", person.firstName, person.lastName)
	fmt.Println()
}

func displayValue(a int) {
	fmt.Println("value of a in deferred function", a)
}

// practical use of defer
type Rectangular struct {
	length int
	width  int
}

func (rectangular Rectangular) area(waitGroup *sync.WaitGroup) int {

	if rectangular.length < 0 {
		fmt.Printf("Rectangular %v's length should be greater than zero\n", rectangular)
		waitGroup.Done()
		return 0
	}

	if rectangular.width < 0 {
		fmt.Printf("Rectangular %v's width should be greater than zero\n", rectangular)
		waitGroup.Done()
		return 0
	}
	waitGroup.Done()
	return (rectangular.length * rectangular.width)
}

//error practice

func openFile() {
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

// Converting the error to the underlying type and retrieving more information from the struct fields

func convertToUndelyingTypeError() {
	file, error := os.Open("test.txt")
	fmt.Println("Your Error is: ", error)
	if error != nil {
		var pathError *os.PathError
		if errors.As(error, &pathError) {
			fmt.Println("Failed to open file at path", pathError.Path)
			return
		}
		fmt.Println("Generic error", error)
		return
	}
	fmt.Println(file.Name(), " opened successfully")
}

func dnsErrorCheck() {
	domain, error := net.LookupHost("golangbot123.com")
	if error != nil {
		var dnsError *net.DNSError
		if errors.As(error, &dnsError) {
			if dnsError.Timeout() {
				fmt.Println("Operation Timeout")
				return
			}

			if dnsError.Temporary() {
				fmt.Println("Temporary Error")
				return
			}

			fmt.Println("Generic DNS error", error)
			return
		}
		fmt.Println("Generic error", error)
		return
	}
	fmt.Println(domain)
}
