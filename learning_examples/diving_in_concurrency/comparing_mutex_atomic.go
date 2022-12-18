package diving_in_concurrency

import (
	"sync"
	"sync/atomic"
	"time"
)

func MutexCounter() int {
	goroutinesCounter := 0
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			goroutinesCounter++
			mu.Unlock()
			time.Sleep(time.Microsecond)
			mu.Lock()
			goroutinesCounter--
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	return goroutinesCounter
}

func AtomicCounter() int32 {
	goroutinesCounter := int32(0)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&goroutinesCounter, 1)
			time.Sleep(time.Microsecond)
			atomic.AddInt32(&goroutinesCounter, -1)

			wg.Done()
		}()
	}

	wg.Wait()
	return goroutinesCounter
}