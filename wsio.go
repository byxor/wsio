// Package wsio provides helpers to hide gorilla websocket connections
// behind the io.Reader and io.Writer interfaces.
package wsio

import (
	"github.com/gorilla/websocket"
)

// WebsocketReader is an io.Reader implementation for gorilla websockets.
type WebsocketReader struct {
	Connection *websocket.Conn
}

// Read reads data from the websocket into the provided slice.
// This function will block further execution until data is received via the websocket.
// Returns the numberOfBytesRead and any errors.
func (self *WebsocketReader) Read(into []byte) (int, error) {
	_, data, err := self.Connection.ReadMessage()
	copy(into, data)
	return len(data), err
}

// WebsocketWriter is an io.Writer implementation for gorilla websockets.
type WebsocketWriter struct {
	Connection *websocket.Conn
}

// Write writes the provided data to the websocket.
// Returns the numberOfBytesWritten and any errors.
func (self *WebsocketWriter) Write(data []byte) (int, error) {
	err := self.Connection.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return 0, err
	}
	return len(data), nil
}
