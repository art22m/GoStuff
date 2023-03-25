package patterns

import (
	"fmt"
	"time"
)

func add(done <-chan struct{}, stream <-chan int, delta int) <-chan int {
	outChan := make(chan int)
	go func() {
		defer close(outChan)
		for v := range stream {
			time.Sleep(1 * time.Second) // Долго, хотель бы распараллелить -> FanIn, FanOut
			select {
			case <-done:
				return

			case outChan <- v + delta:
			}
		}
	}()

	return outChan
}

func pipeline() {
	stream := make(chan int, 5)
	for i := 0; i < 5; i++ {
		stream <- i
	}
	close(stream)

	done := make(chan struct{})
	defer close(done)

	pipeline := add(done, add(done, stream, 2), 1)
	for v := range pipeline {
		fmt.Println(v)
	}
}
