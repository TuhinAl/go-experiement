package main

import (
	"fmt"

	"github.com/pborman/uuid"
	"github.com/tuhinal/go-experiment/concurrency/gochannel"
	"github.com/tuhinal/go-experiment/go-basic/arrayslicemap"
	"github.com/tuhinal/go-experiment/go-basic/cryptit/decrypt"
	"github.com/tuhinal/go-experiment/go-basic/cryptit/encrypt"
)

func main() {
uuid := uuid.NewRandom()
fmt.Println(uuid);
arrayslicemap.Slice();
myEncryptedName := encrypt.Encrypt("Tuhin")
myDecryptedName := decrypt.Decrypt(myEncryptedName)
fmt.Println(myEncryptedName)
fmt.Println(myDecryptedName)
myEncryptedUpperName := gochannel.ToLowercase()
myDecryptedUpperName := gochannel.ToUppercase()
fmt.Println(myEncryptedUpperName)
fmt.Println(myDecryptedUpperName)

}
