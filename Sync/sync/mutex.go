package channels

import (
	"fmt"
	"sync"
)

func mutexTest1() {
	var mu sync.RWMutex
	var cnter int

	go func() {
		mu.RLock()
		defer mu.RUnlock()

		fmt.Println(cnter) // не блокирует на чтение
	}()
}

func mutexTest2() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var cnter int32
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer mu.Unlock()

			mu.Lock()
			cnter++
			//mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(cnter)
}
