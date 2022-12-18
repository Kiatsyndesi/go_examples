package diving_in_concurrency

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cs := map[string]int{"Касса #1": 0, "Касса #2": 0}
	mu := &sync.Mutex{}

	for buyer := 0; buyer < 1000; buyer++ {
		go func(buyer int) {
			mu.Lock()
			defer mu.Unlock()
			cs["Касса #1"] += 1
			fmt.Println(cs)
		}(buyer)
	}

	for buyer := 0; buyer < 1000; buyer++ {
		go func(buyer int) {
			mu.Lock()
			defer mu.Unlock()
			cs["Касса #2"] += 1
			fmt.Println(cs)
		}(buyer)
	}

	time.Sleep(time.Second)
	fmt.Println(cs)
}
