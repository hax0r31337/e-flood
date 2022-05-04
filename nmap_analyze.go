package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var Client = &http.Client{
	Timeout: 7 * time.Second,
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
}

func analyze(address string) {
	ports := ReadLines("res.txt")
	for _, port := range ports {
		go scan(address, strings.Split(port, "/")[0])
		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(5 * time.Second)
}

func scan(address, port string) { // scan a port is websocket
	req, err := http.NewRequest("GET", "http://"+address+":"+port, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Upgrade", "websocket")
	req.Header.Set("Connection", "Upgrade")
	req.Header.Set("Sec-WebSocket-Key", "dGhlIHNhbXBsZSBub25jZQ==")
	req.Header.Set("Sec-WebSocket-Version", "13")
	resp, err := Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(resp.StatusCode)
	if resp.StatusCode == 101 {
		fmt.Println("websocket -> ", port)
		scanmotd(address, port)
	}
}

func scanmotd(address, port string) {
	c, _, err := dial.Dial("ws://"+address+":"+port+"/", nil)
	if err != nil {
		fmt.Println(port+" dial:", err)
		return
	}
	defer c.Close()
	err = c.WriteMessage(websocket.TextMessage, []byte("Accept: MOTD"))
	if err != nil {
		fmt.Println(port+" write:", err)
		return
	}
	_, msg, err := c.ReadMessage()
	if err != nil {
		fmt.Println(port+" read:", err)
		return
	}
	fmt.Println(string(msg))
}
