package simple_goroutines

type Result struct {
	ID     int
	Result int
}

func Consumer(ID int, in <-chan int, out chan<- Result) {
	res := 0

	for value := range in {
		res += value
	}

	out <- Result{ID: ID, Result: res}
}

func main() {

}
