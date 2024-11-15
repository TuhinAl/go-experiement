package main

import (
	"fmt"
	"time"

	"github.com/tuhinal/go-experiment/concurrency/goroutine"
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

	start := time.Now()
	for i := 1; i <= 10000; i++ {
		go goroutine.CalculateSquare(i)
	}
	elapsed := time.Since(start)
	time.Sleep(2 * time.Second)
	fmt.Println("Function took: ", elapsed)

}
