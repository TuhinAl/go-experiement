package selectstatement

import (
	"fmt"
	"sync"
)

// this is a Critical section as I will learn "how to solve race conditions using mutexes and channels."

// Program with a race condition
var x = 0

func increment(wg *sync.WaitGroup) {
	x += 1
	wg.Done()
}

func CreatingRaceCondition() {

	var wg sync.WaitGroup
	// var mutex sync.Mutex
	ch := make(chan bool, 1)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		// go increment(&wg)
		// go incrementWithMutex(&wg, &mutex)
		go incrementWithChannel(&wg, ch)
	}
	wg.Wait()
	fmt.Println("final value of x", x)
}

// Solving the race condition using a mutex

func incrementWithMutex(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	x += 1
	mutex.Unlock()
	wg.Done()
}

//Solving the race condition using channel
func incrementWithChannel(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x += 1
	<-ch
	wg.Done()
}