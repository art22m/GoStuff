package Sync

import (
	"fmt"
	"sync"
)

func wgTest1() {
	mp := map[string]int{}
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		mp["123"] = 1
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		mp["123"] = 2
	}()

	wg.Wait()

	fmt.Println(mp["123"])
}
