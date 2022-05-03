package main

import "time"

func main() {
	LoadWord()
	for true {
		go WsFlood(false)
		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(1000 * time.Second)
}
