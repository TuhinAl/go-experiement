package functions_test

import "fmt"

/*
A closure is a function that captures and "closes over" variables from its surrounding scope.

When a function use a variable out of its scope and can access it, this is called closures

*/
func Fibonacci() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return b
	}
}
func Fib() {
	f, g := Fibonacci(), Fibonacci()
	/* for x := f(); x < 100; x = f(){
		fmt.Println(x)
	} */

	fmt.Println(f(), g(), f(), g(), f())
	fmt.Println(g(), f(), g(), f(), g())
}

func do(d func()) { // function as a parameter(a function that takes no argument and return nothig)
	d()
}

func ClouserDemo() {
	for i := 0; i < 4; i++ {
		v := func() { // The function v captures the variable i. However, it doesn't capture the value of i; it captures the `reference` to i.
			fmt.Printf("%d @ %p\n", i, &i) 
		}
		do(v) // it accesses the current value of i at the time of execution.
	}
}

func ClouserWithSlice() {
	s := make([]func(), 4)

	for i := 0; i < 4; i++ {
		i2 := i //closure capture
		s[i] = func() {
			fmt.Printf("%d @ %p\n", i2, &i2)
		}
	}
	for i := 0; i < 4; i++ {
		s[i]()
	}
}
