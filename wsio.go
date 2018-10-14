package wsio

import (
	"github.com/gorilla/websocket"
)

type WebsocketReader struct {
	Connection *websocket.Conn
}

func (self *WebsocketReader) Read(into []byte) (int, error) {
	_, data, err := self.Connection.ReadMessage()
	copy(into, data)
	return len(data), err
}

type WebsocketWriter struct {
	Connection *websocket.Conn
}

func (self *WebsocketWriter) Write(data []byte) (int, error) {
	err := self.Connection.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return 0, err
	}
	return len(data), nil
}
