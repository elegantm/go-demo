package channel

import "fmt"

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
