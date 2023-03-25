package patterns

import (
	"fmt"
	"sync"
	"time"
)

func donePattern1() {
	done := make(chan struct{}, 2)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			fmt.Println("Iteration..")
			select {
			case <-done:
				fmt.Println("The end")
				return
			}
		}
	}()

	time.Sleep(1 * time.Second)
	close(done)
	wg.Wait()
}

func donePattern2() {
	c := make(chan int) // if we add buffer, and then c <- 1, it will read 1
	//c <- 1 // block
	close(c)

	// Reading from a closed channel
	v, ok := <-c
	fmt.Println(v, ok) // Output: 0 false
}
