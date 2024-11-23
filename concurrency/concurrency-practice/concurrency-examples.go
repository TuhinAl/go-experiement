package concurrencypractice

import (
	"fmt"
	"sync"
)

func GoroutineLeak()  {
	var wg sync.WaitGroup
	wg.Add(2)
	go leak(&wg)
	wg.Wait()
}

func leak(wg *sync.WaitGroup)  {
	ch := make(chan int)
	// ekhane anonymous function ta block hoye ache due to 'waiting for, writing data on ch channel' 
	go func(){
		val := <- ch
		fmt.Println("Received ", val)
		wg.Done()
	}()
	fmt.Println("Exiting Leak method......")
	wg.Done()
}