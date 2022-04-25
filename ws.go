package anet

import (
	"github.com/gorilla/websocket"
)

// WSWrap wrap for websocket
type WSWrap struct {
	c *websocket.Conn
}

// NewWSWrap create wrap
func NewWSWrap(c *websocket.Conn) *WSWrap {
	return &WSWrap{c}
}

func (w *WSWrap) Read(p []byte) (int, error) {
	_, data, err := w.c.ReadMessage()
	if err != nil {
		return 0, err
	}
	return copy(p, data), nil
}

func (w *WSWrap) Write(p []byte) (int, error) {
	return len(p), w.c.WriteMessage(websocket.TextMessage, p)
}
