package main

import (
	"github.com/gorilla/websocket"
)

func WsStuck(address string) {
	c, _, err := dial.Dial("ws://"+address+"/", nil)
	if err != nil {
		// log.Println("dial:", err)
		return
	}
	defer c.Close()
	err = c.WriteMessage(websocket.BinaryMessage, []byte{0x02, 69, 0x00, 0x0d})
	if err != nil {
		// log.Println("write:", err)
		return
	}
	_, _, err = c.ReadMessage()
	if err != nil {
		// log.Println("read:", err)
		return
	}
	// log.Printf("recv: %s", msg)
}

func WsSuperStuck(address string) {
	for true {
		WsStuck(address)
	}
}
