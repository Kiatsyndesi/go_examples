package simple_goroutine

import "fmt"

func producer(ch chan<- string) {
	//Sending to channel
}

func main() {
	ch := make(chan string)

	for i := 0; i < 10; i++ {
		go producer(ch)
	}

	for value := range ch {
		//handling data
		fmt.Print(value)
	}
}
