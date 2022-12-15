package simple_goroutines

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	dataChan := make(chan int)

	for {
		select {
		case <-c:
			fmt.Println("Received interrupted signal")
			return
		case data, ok := <-dataChan:
			if !ok {
				fmt.Println("Something went wrong")
			}
			//handling data
			fmt.Println(data)
		}
	}
}
