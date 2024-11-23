package concurrencypractice

import (
	"fmt"
	"sync"
	"time"
)

// Spawning go-routine closures in a loop
//practice on: Spawning up goroutine inside a closure

func SpawningUpGoroutine() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Done ")
}

//timeout

func Timeout() {
	ch := make(chan int)
	go sendValue(ch)

	select {
	case message := <-ch:
		fmt.Println(message)
	case <-time.After(1 * time.Second):
		fmt.Println("select timeout.")
	}
}

func sendValue(ch chan int) {
	ch <- 15
}
