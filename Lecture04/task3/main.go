package main

import (
	"fmt"
	"sync"
	"time"
)

var cache = make(map[string]string)
var rwMutex sync.RWMutex

func set(key string, value string) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	cache[key] = value
}

func get(key string) string {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	return cache[key]
}

func main() {
	go func() {
		for {
			set("key", "value")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			fmt.Println(get("key"))
			time.Sleep(100 * time.Millisecond)
		}
	}()

	select {}
}
