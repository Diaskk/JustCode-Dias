package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	m    map[string]interface{}
	lock sync.Mutex
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]interface{}),
	}
}

func (sm *SafeMap) Set(key string, value interface{}) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	sm.m[key] = value
}

func (sm *SafeMap) Get(key string) (interface{}, bool) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	val, ok := sm.m[key]
	return val, ok
}

func main() {
	sm := NewSafeMap()
	sm.Set("key", "value")
	value, _ := sm.Get("key")
	fmt.Println(value)
}
