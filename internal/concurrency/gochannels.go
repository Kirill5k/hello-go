package concurrency

import (
	"fmt"
	"time"
)

func GoChannels() {
	const count = 3
	ch := make(chan int)

	go func() {
		for i := 0; i < count; i++ {
			fmt.Printf("sending %d over channel\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	/*
		for i := 0; i < count; i++ {
			val := <-ch
			fmt.Printf("Received %d from channel\n", val)
		}
	*/

	for val := range ch {
		fmt.Printf("Received %d from channel\n", val)
	}
}
