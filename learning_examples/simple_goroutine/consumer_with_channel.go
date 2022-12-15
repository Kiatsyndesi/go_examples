package simple_goroutine

func consumer(ch <-chan string) {
	//Receiving from channel
}

func main() {
	ch := make(chan string)

	for i := 0; i < 10; i++ {
		go consumer(ch)
	}

	for value := range ch {
		ch <- value
	}
}
