package goroutine

import (
	"fmt"
	"sync"
)
func CalculateSquareUsingWaitGroup(i int, wg *sync.WaitGroup) {
	fmt.Println(i * i)
	wg.Done();
}