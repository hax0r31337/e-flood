package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	LoadWord()
	for true {
		go WsFlood(false)
		time.Sleep(100 * time.Millisecond)
	}
	// go WsFlood(false)
	time.Sleep(1000 * time.Second)
}
