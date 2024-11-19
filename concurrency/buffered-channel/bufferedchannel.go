package bufferedchannel

import (
	"fmt"
	"sync"
)

func Bufferedchannel() {
	var wg sync.WaitGroup
	wg.Add(2)               // since I'm plainning to create 2 gorouitne
	ch := make(chan int, 3) // I have created a channel that can transfer only int values and has a buffer of 3
	go sell1(ch, &wg)
	wg.Wait()
}

func sell1(ch chan int, wg *sync.WaitGroup) {
	ch <- 10
	ch <- 11
	ch <- 12
	go buy1(ch, wg)
	ch <- 13  // for buffered channel (which send operation) block the goroutine if it exceed the buffer limit
	
	fmt.Println("Sent all the data to the channel")
	wg.Done()
}

func buy1(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Waiting for the data .....")
	fmt.Println("Received data: ", <-ch)
	wg.Done()
}


