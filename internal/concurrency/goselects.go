package concurrency

import (
	"fmt"
	"time"
)

func GoSelect() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		ch1 <- 42
	}()

	select {
	case val := <-ch1:
		fmt.Printf("Received value from ch1: %d\n", val)
	case val := <-ch2:
		fmt.Printf("Received value from ch2: %d\n", val)
	}

	chOut := make(chan float64)
	go func() {
		time.Sleep(100 * time.Millisecond)
		chOut <- 3.14
	}()

	select {
	case val := <-chOut:
		fmt.Printf("Received %f\n", val)
	case <-time.After(20 * time.Millisecond):
		fmt.Println("timeout")
	}
}
