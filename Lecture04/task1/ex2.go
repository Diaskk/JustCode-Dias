package main

import (
	"fmt"
	"time"
)

var balances = make(map[string]int)

func updateBalance(account string, amount int) {
	balance := balances[account]
	balance += amount
	balances[account] = balance
}

func main() {
	balances["Alice"] = 100
	go updateBalance("Alice", 50)
	go updateBalance("Alice", -30)
	time.Sleep(time.Second)
	fmt.Println(balances["Alice"])
}
