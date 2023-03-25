package patterns

import (
	"fmt"
	"time"
)

func startTimer() {
	timer := time.NewTimer(2 * time.Second)

	for {
		select {
		case <-timer.C:
			fmt.Println("end")
			return

		/*
			time.After() function returns a new channel every time it is called,
			which means that the channel being used in the first case of the select statement
			is not the same channel as the one being checked in the loop.
			As a result, the timer in the first case is not being reset every time through the loop,
			and so the loop never exits.
		*/
		//case <-time.After(2 * time.Second):
		//	fmt.Println("end")
		//	return
		default:
		}

		fmt.Println("test")
		time.Sleep(1 * time.Second)
	}
}
