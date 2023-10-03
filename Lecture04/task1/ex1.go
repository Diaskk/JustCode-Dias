package main

import (
	"fmt"
	"time"
)

var counter int

func increment() {
	counter++
}

func main() {
	go increment()
	go increment()
	time.Sleep(time.Second)
	fmt.Println(counter)
}
