package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/tuhinal/go-experiment/concurrency/goroutine"
	"github.com/tuhinal/go-experiment/concurrency/buffered-channel"
)

func main() {
	/* uuid := uuid.NewRandom()
	   fmt.Println(uuid);
	   arrayslicemap.Slice();
	   myEncryptedName := encrypt.Encrypt("Tuhin")
	   myDecryptedName := decrypt.Decrypt(myEncryptedName)
	   fmt.Println(myEncryptedName)
	   fmt.Println(myDecryptedName)
	   myEncryptedUpperName := gochannel.ToLowercase()
	   myDecryptedUpperName := gochannel.ToUppercase()
	   fmt.Println(myEncryptedUpperName)
	   fmt.Println(myDecryptedUpperName) */

	/* start := time.Now()
	for i := 1; i <= 10000; i++ {
		go goroutine.CalculateSquare(i)
	}
	elapsed := time.Since(start)
	time.Sleep(2 * time.Second)
	fmt.Println("Function took: ", elapsed) */
	var wg sync.WaitGroup
	start := time.Now()
	wg.Add(10000)
	for i := 1; i <= 10000; i++ {
		go goroutine.CalculateSquareUsingWaitGroup(i, &wg)
	}
	elapsed := time.Since(start)
	wg.Wait()
	fmt.Println("Function toook: ", elapsed)

}
