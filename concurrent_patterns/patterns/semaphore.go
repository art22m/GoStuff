package patterns

import (
	"fmt"
	"time"
)

type semaphore struct {
	semC chan struct{}
}

func (s *semaphore) Acquire() {
	s.semC <- struct{}{}
}

func (s *semaphore) Release() {
	<-s.semC
}

func semaphoreTest() {
	const limitGoroutines = 2

	s := &semaphore{
		semC: make(chan struct{}, limitGoroutines),
	}

	for i := 0; i < 10; i++ {
		s.Acquire()
		go func(i int) {
			defer s.Release()
			fmt.Println(i)
			time.Sleep(2 * time.Second)
		}(i)
	}
}
