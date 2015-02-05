package hub

import (
	"io"
)

type hub struct {
	connections map[io.WriteCloser]bool
}

type Hub interface {
	Register(client io.WriteCloser)
	Unregister(client io.WriteCloser)
	Broadcast(message []byte)
}

func NewHub() *hub {
	return &hub{
		connections: make(map[io.WriteCloser]bool),
	}
}

func (h *hub) Register(c io.WriteCloser) {
	h.connections[c] = true
}

func (h *hub) Unregister(c io.WriteCloser) error {
	if _, ok := h.connections[c]; ok {
		delete(h.connections, c)
		return c.Close()
	}
	return nil
}

func (h *hub) Broadcast(m []byte) {
	for c := range h.connections {
		if _, err := c.Write(m); err != nil {
			h.Unregister(c)
		}
	}
}
