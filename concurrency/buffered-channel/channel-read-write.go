package bufferedchannel

import (
	"fmt"
	"time"
)

func ChannelRead() {
	ch := make(chan string)
	go buy(ch)
	go sell(ch)
	time.Sleep(1 * time.Second)
}

// send data to the channel
func sell(ch chan string) {
	ch <- "Real Madrid"
	fmt.Println("Send data to the channel")
}

// received data from channel
func buy(ch chan string) {
	fmt.Println("Waiting for the data")
	data := <-ch
	fmt.Println("Received data.... ", data)
}
