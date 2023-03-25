package main

import (
	"fmt"
	"time"
)

func main() {

}

// TEST 3

func myFun(in chan<- int) {
	in <- 1
	fmt.Println(1)
	in <- 2
	fmt.Println(2)
	in <- 3
	fmt.Println(3)
	in <- 4
	fmt.Println(4)
	in <- 5
	fmt.Println(5)
}

func chanTest3() {
	ch := make(chan int, 3) // 1 2 3 ... 1 4
	go myFun(ch)
	time.Sleep(1 * time.Second)
	fmt.Println(<-ch)
	time.Sleep(1 * time.Second)
}

// TEST 2

func printer(in <-chan int) {
	fmt.Println("printer")
	fmt.Println(<-in)
}

func sender(in chan<- int) {
	fmt.Println("sender")
	in <- 2
}

func chanTest2() {
	ch := make(chan int)
	go sender(ch)
	time.Sleep(1.0 * time.Second)
	go printer(ch)
	time.Sleep(1.0 * time.Second)
}

// TEST 1

func chanTest1() {
	ch := make(chan int)

	go func() {
		ch <- 12
	}()

	// VAR 1

	//val, _ := <-ch // ok -- состояние канала (закрыт / открыт)

	// VAR 2

	//for v := range ch {
	//	fmt.Println(v)
	//}

	// VAR 3

	//select {
	//case val, ok := <-ch:
	//	if ok {
	//		fmt.Println(val)
	//	}
	//
	//default:
	//	fmt.Println("Cannot read value")
	//}

	close(ch)
}
