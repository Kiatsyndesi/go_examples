package simple_goroutines

import (
	"fmt"
	"runtime"
	"time"
)

func outputFiveByFiveTimes() {
	for i := 0; i < 5; i++ {
		go func() {
			time.Sleep(time.Second)
			fmt.Println(i)
		}()
	}

	time.Sleep(time.Second * 2)
}

func outputRandomNumbersFiveTimes() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			time.Sleep(time.Second)
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second * 2)
}

func outputOneByOne() {
	runtime.GOMAXPROCS(1)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	runtime.Gosched()

	time.Sleep(time.Second)
}

func main() {
	outputFiveByFiveTimes()

	outputRandomNumbersFiveTimes()

	outputOneByOne()
}
