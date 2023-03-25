package patterns

import (
	"fmt"
	"sync"
	"time"
)

func addSlow(done <-chan struct{}, stream <-chan int, delta int) <-chan int {
	outChan := make(chan int)
	go func() {
		defer close(outChan)
		for v := range stream {
			time.Sleep(1000)
			select {
			case outChan <- v + delta:

			case <-done:
				return
			}
		}
	}()

	return outChan
}

func fanIn(done <-chan struct{}, chans ...<-chan int) <-chan int {
	multiplexed := make(chan int)

	var wg sync.WaitGroup
	for _, ch := range chans {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				select {
				case <-done:
					return

				case multiplexed <- v:
				}
			}

		}(ch)
	}

	go func() {
		wg.Wait()
		close(multiplexed)
	}()

	return multiplexed
}

// Will be used from Pipeline file

//func add(done <-chan struct{}, stream <-chan int, delta int) <-chan int {
//	outChan := make(chan int)
//	go func() {
//		defer close(outChan)
//		for v := range stream {
//			time.Sleep(500)
//			select {
//			case outChan <- v + delta:
//
//			case <-done:
//				return
//			}
//		}
//	}()
//
//	return outChan
//}

func fanInFanOut() {
	done := make(chan struct{})
	defer close(done)

	stream := make(chan int, 5)
	fanOut := make([]<-chan int, 0, 5)
	for i := 0; i < 5; i++ {
		stream <- i
		fanOut = append(fanOut, addSlow(done, stream, 1))
	}
	close(stream)

	start := time.Now()
	pipeline := add(done, fanIn(done, fanOut...), 2)
	for v := range pipeline {
		fmt.Println(v)
	}
	fmt.Println(time.Since(start))
}
