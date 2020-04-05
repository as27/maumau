package main

import "github.com/gorilla/websocket"

type client struct {
	socket   *websocket.Conn
	messages chan []byte
}
