package selectstatement

import (
	"fmt"
	"time"
)

func SelectDemo() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	//go goOne(ch1)
	//go goTwo(ch2)
	go server1(ch1)
	go server2(ch2)

	select {
	case val1 := <-ch1:
		fmt.Println(val1)
	case val2 := <-ch2:
		fmt.Println(val2)
	}
	time.Sleep(1 * time.Second)
}

func goTwo(ch2 chan string) {
	ch2 <- "Channel-2"
}

func goOne(ch1 chan string) {
	ch1 <- "Channel-1"
}

func server1(ch2 chan string) {
	time.Sleep(6 * time.Second)
	ch2 <- "from server-1"
}

func server2(ch1 chan string) {
	time.Sleep(3 * time.Second)
	ch1 <- "from server-2"
}

func DefaultCaseOfSelect() {
	ch := make(chan string)
	go process(ch)

	for {
		fmt.Println("wait for a second....")
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received ")
		}
	}

}

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func DeadlockAndDefaultCase() {
	ch := make(chan string)

	select {
	case <-ch:
	default:
		fmt.Println("since no goroutine is writing onto this channel, deadlock will occured, that is why I keep a default case...")
	}
}

func NilChannelDefaultCase() {
	var ch chan string

	select {
	case <-ch:
	default:
		fmt.Println("for nil channel, default case will executed...")
	}
}

func RandomSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go webserver(ch1)
	go DBserver(ch2)
	time.Sleep(1000 * time.Millisecond)
	select {
	case s1 := <-ch1:
		fmt.Println(s1)
	case s2 := <-ch2:
		fmt.Println(s2)
	}

}

func DBserver(ch2 chan string) {
	ch2 <- "This is from DB server."
}

func webserver(ch1 chan string) {
	ch1 <- "This is from Web server."
}
