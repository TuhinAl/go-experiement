package main

import (
	"fmt"

	"github.com/pborman/uuid"
	"github.com/tuhinal/go-experiment/go-basic/cryptit/encrypt"
)

func main() {
uuid := uuid.NewRandom()
fmt.Println(uuid);
//arrayslicemap.Slice();
myName := encrypt.Encrypt("Tuhin")
fmt.Println(myName);
}
