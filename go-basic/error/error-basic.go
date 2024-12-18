package error__test

import (
	"errors"
	"fmt"
	"net"
	"os"
	"reflect"
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

func dividebyZero(a, b int) (int, error) {
	if b == 0 {
		// return 0.00, errors.New("Divide by Zero")
		return 0.00, &IllegalArgumentError{Code: 501, Message: "Internal Server Error"}
	}
	return a / b, nil
}

func PanicInGolang() {
	// BasicDBConfigMissingPanic()
	HandlePanicUsingRecovery()
}

func ReflectionInGolang() {
	ReflectionWithComplexType()
}

func ErrorInGolang() {
	// result, err := dividebyZero(10, 0)

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Result:", result)
	// fmt.Println("==========================")
	ErrorWrapAndUnWrap()
}

type IllegalArgumentError struct {
	Code    int
	Message string
}
type ValidatioError struct {
	Field, Message string
}
type DatabaseConnectionError struct {
	ConnectionHost, Port string
}

func (v *ValidatioError) Error() string {
	return fmt.Sprintf("Field: %s\nMessage: %s\n", v.Field, v.Message)
}

func (d *DatabaseConnectionError) Error() string {
	return fmt.Sprintf("Connection: %s\nPort: %s\n", d.ConnectionHost, d.Port)
}

func (e *IllegalArgumentError) Error() string {
	return fmt.Sprintf("Code: %d\nMessage: %s\n", e.Code, e.Message)
}

func (e *IllegalArgumentError) Is(target error) bool {
	t, ok := target.(*IllegalArgumentError)
	if !ok {
		return false
	}
	return e.Code == t.Code && e.Message == t.Message
}

func ErrorWrapAndUnWrap() {
	fileName := "data.txt"
	err := fileProcessing(fileName)
	if err != nil {
		fmt.Println(err)
	}
	unknownError := &IllegalArgumentError{Code: 503, Message: "Internal"}
	if errors.Is(err, unknownError) {
		fmt.Println("Error is unknown")
	}

	ErrorAs()
}

func fileProcessing(fileName string) error {
	illegalError := &IllegalArgumentError{Code: 503, Message: "Internal"}
	return fmt.Errorf("your file name %s cannot be open %w", fileName, illegalError)
}

func PerformValidation() error {
	//simulate an error and wrap it
	vError := &ValidatioError{Field: "email", Message: "There is no @ sign"}
	return fmt.Errorf("action failed because of validation error: %w", vError)
}

func PerformDBAction() error {
	//simulate an error and wrap it
	return fmt.Errorf("validation error: %w", &DatabaseConnectionError{ConnectionHost: "localhost", Port: "8080"})
}

func ErrorAs() {
	err := PerformValidation()
	err2 := PerformDBAction()
	var dbError *DatabaseConnectionError
	var validationError *ValidatioError
	if errors.As(err, &validationError) {
		fmt.Printf("Validation Error: Field = %s\n", validationError.Field)

	}
	if errors.As(err2, &dbError) {
		fmt.Printf("Database connection Error: host: %s port = %s\n", dbError.ConnectionHost, dbError.Port)

	}
	fmt.Println("Unknown error:", err)
}

func BasicDBConfigMissingPanic() {
	dbConfig := ""
	if dbConfig == "" {
		panic("Database Configuration missing..")
	}
	fmt.Println("Connecting to database.......")
}

func HandlePanicUsingRecovery() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	file, err := os.Open("student.txt")
	if err != nil {
		panic(fmt.Sprintf("failed to opend file: %v", err))
	}
	defer file.Close()
	fmt.Println("File opened successfully!")
}

func basicReflection() {
	data := 40
	fmt.Println("Type: ", reflect.TypeOf(data))
	fmt.Println("Value: ", reflect.ValueOf(data))
}

const (
	_ errKind = iota
	noHeader
	cantReadHeader
	invalidHdrType
	invalidChkLength
)

type errKind int
type WaveError struct {
	kind  errKind
	value int
	err   error
}

func (e WaveError) Error() string {
	switch e.kind {
	case noHeader:
		return "no header (file too short?)"
	case cantReadHeader:
		return fmt.Sprintf("can't read header[%d]: %s", e.value, e.err.Error())
	case invalidHdrType:
		return "invalid header type"
	case invalidChkLength:
		return fmt.Sprintf("invalid chunk length: %d", e.value)

	}

	func handler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world! from %s\n", r.URL.Path[1:])
	}

}
