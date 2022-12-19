package diving_in_concurrency

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cs := map[string]int{"Cash #1": 0, "Cash #2": 0}
	mu := &sync.Mutex{}

	for buyer := 0; buyer < 1000; buyer++ {
		go func(buyer int) {
			mu.Lock()
			//can be without defer, but on this way we should move unlock to end
			defer mu.Unlock()
			cs["Cash #1"] += 1
			fmt.Println(cs)
		}(buyer)
	}

	for buyer := 0; buyer < 1000; buyer++ {
		go func(buyer int) {
			mu.Lock()
			defer mu.Unlock()
			cs["Cash #2"] += 1
			fmt.Println(cs)
		}(buyer)
	}

	time.Sleep(time.Second)
	fmt.Println(cs)
}
