package main

import (
	"fmt"
	"sync"
	"time"
)

func countdown(name string, from int, group *sync.WaitGroup) {
	for i := 0; i < from; i++ {
		fmt.Printf("countdown (%s): %d\n", name, i)
		time.Sleep(time.Second)
	}
	group.Done()
}

func main() {
	// wait groups are a way to wait for multiple goroutines to finish.
	// you tell it how many 'deltas' to wait for. in each goroutine, you call Done()
	// to decrement the internal counter
	var wg sync.WaitGroup
	wg.Add(2)

	// goroutines are created by calling any function with the 'go' keyword.
	go countdown("1", 5, &wg)
	go countdown("2", 3, &wg)

	// finally block until all the groups to finish
	wg.Wait()
}
