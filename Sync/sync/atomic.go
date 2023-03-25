package channels

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func atomicTest1() {
	var wg sync.WaitGroup
	var cnter int32
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cnter++
			atomic.AddInt32(&cnter, 1)
		}()
	}

	wg.Wait()
	fmt.Println(cnter)
}
