package datatypeoperator

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	a = 1
	b = 2 * 1024
	c = b << 3
	s = "A string"
	t = len(s)
	// u = s[2:]    Syntax error
)

/*
Golang types falls into Four category:

 1. Basic Type: numbers, strings, and booleans.

 2. Aggregate Type: Array & structs

 3. Reference Type: Pointer, Slice, Map, Function & Channel

 4. Interface Type

    Four standard packages are particularly importantt for manipulating string:
    bytes, strings, strconv, unicode
*/
const boilingF = 212.0

func main1() {
	fmt.Println("================================= data type =============================")

	/*	there are two types of data type
			1. Basic Type
				(i) Numeric
		 		(ii) String
				(iii) Boolean

			2. Composite Type
	*/

	str := "Welcome to the Go Programming language"
	id := 350
	result := 3.63
	//var res int8;
	boolean := false
	//char := 'T'
	//println("===========String Datatype ======")
	//fmt.Printf("%s\n", str)
	//fmt.Printf("%q\n", str)
	//fmt.Printf("%x\n", str)
	//fmt.Printf("%X\n", str)

	//println("===========int Datatype ======")
	//fmt.Printf("%d\n", id)
	//fmt.Printf("%b\n", id)
	//fmt.Printf("%o\n", id)
	//fmt.Printf("%x\n", id)
	//fmt.Printf("%X\n", id)

	//println("===========float Datatype ======")
	//fmt.Printf("%f\n", result)
	//fmt.Printf("%e\n", result)
	//fmt.Printf("%E\n", result)
	//fmt.Printf("%g\n", result)
	fmt.Printf("%t\n", boolean)

	println("===========printing pointer address ======")
	fmt.Printf("%p\n", &str)
	fmt.Printf("%p\n", &id)
	fmt.Printf("%p\n", &result)
	fmt.Printf("%p\n", &boolean)

	// Signed and unsigned integer

	//Signed means positive and negative numbers { -5 to 5}
	// unsigned means only positive numbers { 0 to 5}

}

func FindDuplicateLine() {

}

func CalculateBoilingPoint() {
	// var fahrenheit = boilingF
	var fahrenheit = 33.81
	// var celcius = (fahrenheit - 32 ) * 5 / 9
	celcius := FahrenheitToCelcius(fahrenheit)
	fmt.Printf("boiling point %g°F or %g°C\n", fahrenheit, celcius)
}

func FahrenheitToCelcius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func HasPrefix(str, prefix string) bool {
	return len(str) >= len(prefix) && str[:len(prefix)] == prefix
}

func HasSuffix(str, suffix string) bool {
	return len(str) >= len(suffix) && str[len(str)-len(suffix):] == suffix
}

func IsSubString(str, substring string) bool {

	for i := 0; i < len(str); i++ {
		if HasPrefix(str[i:], substring) {
			return true
		}
	}
	return false
}

func UnicodeCodePoint() {
	str := "Hello, 你好"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))

	for i := 0; i < len(str); {
		r, size := utf8.DecodeLastRuneInString(str[i:])
		fmt.Println("r is: ", r)
		fmt.Println("size is", size)
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	unicodeStr := "你好"
	stringAsRune := []rune(unicodeStr)

	fmt.Println(len(stringAsRune))
}

func InputFromTerminalAndFile() {

	var sum float64
	var times int
	for {
		var value float64

		_, err := fmt.Fscanln(os.Stdin, &value)

		if err != nil {
			break
		}
		sum += value
		times++
	}
	if times == 0 {
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}

	fmt.Println("The average is: ", sum/float64(times))
}

func Stringreplace() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough args")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, new)
		fmt.Println(t)
	}
}
