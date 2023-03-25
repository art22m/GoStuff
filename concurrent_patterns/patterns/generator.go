package patterns

import "fmt"

func generator(start, end int) <-chan int {
	result := make(chan int)

	go func() {
		for i := start; i < end; i++ {
			result <- i
		}
		close(result)
	}()

	return result
}

func startGenerator() {
	for val := range generator(0, 10) {
		fmt.Println(val)
	}
}
