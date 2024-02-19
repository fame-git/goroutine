package main

import (
	"fmt"
)

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go func() {
		channel1 <- 10
		close(channel1)
	}()

	go func() {
		channel2 <- 20
		close(channel2)
	}()

	closedChannel1, closedChannel2 := false, false

	for {
		if closedChannel1 && closedChannel2 {
			break
		}
		select {
		case v, ok := <-channel1:
			if !ok {
				closedChannel1 = true
				continue
			}
			fmt.Println("Channel 1", v)
		case v, ok := <-channel2:
			if !ok {
				closedChannel2 = true
				continue
			}
			fmt.Println("Channel 2", v)
		}
	}
}
