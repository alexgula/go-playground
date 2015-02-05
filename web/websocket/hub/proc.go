package hub

import (
	"io"
)

type hubChannels struct {
	broadcast  chan []byte
	register   chan io.WriteCloser
	unregister chan io.WriteCloser
}

func NewHubAsync() *hubChannels {
	return &hubChannels{
		broadcast:  make(chan []byte),
		register:   make(chan io.WriteCloser),
		unregister: make(chan io.WriteCloser),
	}
}

func (p *hubChannels) Run(h Hub) {
	for {
		select {
		case c := <-p.register:
			h.Register(c)
		case c := <-p.unregister:
			h.Unregister(c)
		case m := <-p.broadcast:
			h.Broadcast(m)
		}
	}
}

func (h *hubChannels) Register(client io.WriteCloser) {
	h.register <- client
}

func (h *hubChannels) Unregister(client io.WriteCloser) {
	h.unregister <- client
}

func (h *hubChannels) Broadcast(message []byte) {
	h.broadcast <- message
}
