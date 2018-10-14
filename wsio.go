package wsio

import (
	"github.com/gorilla/websocket"
)

type WebsocketReader struct {
	connection *websocket.Conn
}

func (self *WebsocketReader) Read(into []byte) (int, error) {
	_, data := self.connection.ReadMessage()
	copy(into, data)
	return len(data), err
}

type WebsocketWriter struct {
	connection *websocket.Conn
}

func (self *WebsocketWriter) Write(data []byte) (int, error) {
	err := self.connection.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return 0, err
	}
	return len(data), err
}
