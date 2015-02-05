package connection

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type connection struct {
	ws *websocket.Conn

	send chan []byte
}

func (conn *connection) Read() {

}

func (conn *connection) Write(p []byte) (n int, err error) {
	conn.send <- p
}

func (conn *connection) Close() error {
	close(conn.send)
}
