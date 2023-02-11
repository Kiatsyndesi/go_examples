package diving_in_concurrency

import (
	"context"
	"fmt"
	"sync/atomic"
)

func main() {
	//marshalling example with addint32
	value1 := int32(0)
	atomic.AddInt32(&value1, 1)
	fmt.Println(value1)

	//second example with atomic interface
	value2 := atomic.Value{}
	value2.Store(100)
	ct := value2.Load()
	fmt.Println(ct)

	//third example with context: lose performance
	ctx := context.Background()
	value3 := atomic.Value{}
	value3.Store(&ctx)
	ctFromThird := value3.Load().(context.Context)
	fmt.Println(ctFromThird)
}
