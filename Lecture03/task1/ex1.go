package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1 // попытка отправить данные в канал но никто не принимает
	}()
	time.Sleep(time.Second) // ждем чтобы горутина начала выполнение
	<-ch                    // пытаемся принять данные но они уже отправлены
	fmt.Println("Main function completed")
}
