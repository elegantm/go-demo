package channel

import (
	"fmt"
	"time"
)

// read channel with close
// result : panic
func ReadCloseChannel() {
	ch := make(chan int, 3)
	close(ch)
	ch <- 1
}

func ReadCloseChannelWithData() {
	ch := make(chan float32, 3)
	ch <- 1.11
	ch <- 1.233
	close(ch)

	a, ok := <-ch
	fmt.Println(a, ok)

	a, ok = <-ch
	fmt.Println(a, ok)

	a, ok = <-ch
	fmt.Println("read data:", a, ok)

}

func ChanAndArray() {
	//ch := make(chan int, 1)
	var ch chan int

	fmt.Println("ptr: ", ch)
	fmt.Println("ptr ", &ch)
	fmt.Println("is nil ", ch == nil)

	ch <- 3

	time.Sleep(1 * time.Second)

}
