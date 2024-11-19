package bufferedchannel

import "fmt"

func ChannelClose() {
	ch := make(chan int, 10)
	ch <- 10
	ch <- 11
	ch <- 12
	val, ok := <-ch
	fmt.Println(val, ok)
	close(ch)
	val, ok = <-ch
	fmt.Println(val, ok)

	val, ok = <-ch
	fmt.Println(val, ok)

	val, ok = <-ch
	fmt.Println(val, ok)

	val, ok = <-ch
	fmt.Println(val, ok)
}

// check how panic situatioan arise

func Panic() {
	ch := make(chan int, 10)
	ch <- 10
	ch <- 11

	val, ok := <-ch
	fmt.Println(val, ok)
	close(ch)
	close(ch)
	ch <- 12 //sending to a channel after it has been closed. It will arise Panic situation
	ch <- 13 

	val, ok = <-ch
	fmt.Println(val, ok)

	val, ok = <-ch
	fmt.Println(val, ok)

	val, ok = <-ch
	fmt.Println(val, ok)
}


func ChannelForRange()  {
	nums := [10]int{32, 15, 53, 21, 63, 56}
	for index, value := range  nums{
		fmt.Println("index- ",index,"value- ", value)
	}
}