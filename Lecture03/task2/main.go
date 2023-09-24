package main

import (
	"fmt"
	"sync"
)

func merge(chans ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	for _, ch := range chans {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				out <- n
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()

	go func() {
		ch2 <- 3
		ch2 <- 4
		close(ch2)
	}()

	for v := range merge(ch1, ch2) {
		fmt.Println(v)
	}
}
