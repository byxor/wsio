// The wsio package provides io.Reader & io.Writer wrappers for Gorilla Websockets.
package wsio

import (
	"github.com/gorilla/websocket"
)

// The WebsocketReader type implements io.Reader.
type WebsocketReader struct {
	Connection *websocket.Conn
}

// The `Read` function reads data from the websocket.
// The data is read into the `into` buffer.
// Returns the `numberOfBytesRead` and any `errors`.
// Note: This function will block execution until data is received.
func (self *WebsocketReader) Read(into []byte) (int, error) {
	_, data, err := self.Connection.ReadMessage()
	copy(into, data)
	return len(data), err
}

// The WebsocketWriter type implements io.Writer.
type WebsocketWriter struct {
	Connection *websocket.Conn
}

// The `Write` function writes data to the websocket.
// Pass your data as the first parameter.
// Returns the `numberOfBytesWritten` and any `errors`.
func (self *WebsocketWriter) Write(data []byte) (int, error) {
	err := self.Connection.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return 0, err
	}
	return len(data), nil
}
