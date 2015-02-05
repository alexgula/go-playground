package main

import (
	"github.com/gorilla/websocket"
	//"net/http"
)

type connection struct {
	ws *websocket.Conn

	send chan []byte
}
