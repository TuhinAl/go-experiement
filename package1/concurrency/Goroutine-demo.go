package main

import (
	"fmt"
)

func main() {
	//done := make(chan bool) // declare channel
	//go hello(done)
	//<-done // read data from channel // here blocking happens
	////time.Sleep(1 * time.Second)
	//fmt.Println("main function")

	// better concurrency
	number := 589
	squareChannel := make(chan int)
	cubeChannel := make(chan int)
	go calculateSquares(number, squareChannel)
	go calculateCube(number, cubeChannel)
	squares, cubes := <-squareChannel, <-cubeChannel
	fmt.Println("Final output of squares and cubes:", squares+cubes)

	//deadlock happen
	/*ch := make(chan int)
	ch <- 5*/
}

func hello(done chan bool) {

	fmt.Println("Hello world goroutine")
	done <- true // write data to the channel
}

func calculateSquares(number int, squareChannel chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareChannel <- sum
}

func calculateCube(number int, cubeChannel chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeChannel <- sum
}
