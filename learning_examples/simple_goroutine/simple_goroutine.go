package simple_goroutine

import (
	"fmt"
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

func main() {
	outputFiveByFiveTimes()

	outputRandomNumbersFiveTimes()

}
