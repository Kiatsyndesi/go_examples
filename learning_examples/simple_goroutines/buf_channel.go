package simple_goroutines

import (
	"fmt"
	"time"
)

func worker(ch <-chan int) {
	for {
		v, ok := <-ch

		if !ok {
			return
		}

		fmt.Println(v)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int, 5)

	for i := 0; i < 100; i++ {
		ch <- i
		go worker(ch)
	}
}
