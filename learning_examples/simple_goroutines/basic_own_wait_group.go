package simple_goroutines

import "fmt"

type Wait struct {
	count int
	ch    chan interface{}
}

func InitWait(count int) *Wait {
	return &Wait{count: count, ch: make(chan interface{}, count)}
}

func (w *Wait) Done() {
	w.ch <- 1
}

func (w *Wait) Wait() {
	defer close(w.ch)

	for {
		<-w.ch
		w.count-- // not atomic operation
		if w.count <= 0 {
			return
		}
	}
}

func test(id int, w *Wait) {
	defer w.Done()
	fmt.Printf("%d done", id)
}

func main() {
	wg := InitWait(5)

	for i := 0; i < 5; i++ {
		test(i, wg)
	}

	wg.Wait()
	fmt.Println("Done")
}
