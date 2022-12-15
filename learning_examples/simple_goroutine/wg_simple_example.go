package simple_goroutine

import (
	"fmt"
	"sync"
)

func longOperation(i int, wg *sync.WaitGroup) {
	wg.Done()
	fmt.Print(i)
}

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go longOperation(i, wg)
	}

	wg.Wait()
}
