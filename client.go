package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type client struct {
	socket   *websocket.Conn
	messages chan []byte
	playerID string
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.messages {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("error client.write:", err)
			return
		}
	}
}
