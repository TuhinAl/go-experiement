package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go process(i, &waitGroup)
	}

	waitGroup.Wait()
	fmt.Println("All go routines finished executing")
}

func process(index int, waitGroup *sync.WaitGroup) {
	fmt.Println("started Goroutine ", index)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", index)
	waitGroup.Done()
}
