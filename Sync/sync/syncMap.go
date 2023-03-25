package channels

import "sync"

func smTest1() {
	var mp sync.Map

	go func() {
		mp.Store("123", "13")
	}()

	go func() {
		mp.Store("1fds", "13")
	}()
}
