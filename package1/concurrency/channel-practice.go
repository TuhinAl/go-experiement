package main

import (
	"fmt"
	"time"
)

/**
channel: Channels can be thought of as pipes using which Goroutines communicate. Similar to how water flows from one end
to another in a pipe, data can be sent from one end and received from the other end using channels.
*/

func main7() {
	//channelCreate()

	// Blocking Goroutine and Channel
	//declare channel in this part
	//done := make(chan bool)
	//value := make(chan int)
	ch := make(chan int, 2)
	/*go Hello(done)
	//time.Sleep(1 * time.Second)
	//Legal: read/receiving from channel and unassign
	<-done
	fmt.Println("main function")*/

	// Better version Hello()
	/*fmt.Println("Main going to call hello go goroutine")
	go betterVersionHello(done)
	<-done
	fmt.Println("Main received data")*/
	//Deadlock
	//readDeadlock()
	//writeDeadlock()

	//Unidirectional channel
	// send to channel
	//go sendChannel(value)
	//go receiveChannel(value)
	//time.Sleep(time.Second * 1)

	//channel closing using range loop
	/*	go producePrimeNumber(value)
		for {
			value, ok := <-value
			if !ok {
				break
			}
			fmt.Println("Received ", value, ok)
		}*/

	//unBufferedChannelWrite
	go unBufferedChannelWrite(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)

	}
}

func channelCreate() {
	var channel1 chan int
	if channel1 == nil {
		fmt.Println("channel a is nil, going to define it")
		channel1 = make(chan int)
		fmt.Printf("Type of a is %T", channel1)
	}
}

/*
*
Sends and receives to a channel are blocking by default
*/
func sendingAndReceivingChannel() {
	var channel1 chan int

	readFromChannel1 := <-channel1
	fmt.Println("read from a channel ", readFromChannel1)
	value := 50
	channel1 <- value
	fmt.Println("write to a channel ", <-channel1)
}

func Hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	//write to the channel
	done <- true
}

func betterVersionHello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}
func readDeadlock() {
	ch := make(chan int)
	ch <- 5
}
func writeDeadlock() {
	ch := make(chan int)
	<-ch
}

/**
Unidirectional channels: a channel can send or receive data
*/

func sendChannel(sendCh chan<- int) {
	data := 15
	sendCh <- data
	fmt.Println("Send Success:")
}
func receiveChannel(readCh <-chan int) {

	data := <-readCh
	fmt.Println("Received:", data)
}

func producePrimeNumber(channel chan int) {
	for i := 6; i <= 200; i++ {
		if i%2 != 0 && i%3 != 0 && i%4 != 0 && i%5 != 0 && i%6 != 0 && i%7 != 0 {
			channel <- i
		}
	}
	close(channel)
}

func unBufferedChannelWrite(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
