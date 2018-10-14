package main

import (
	"fmt"
	"github.com/byxor/wsio"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
)

const (
	route      = "/"
	address    = "localhost:8080"
	bufferSize = 64000
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc(route, echoHandler)
	log.Fatal(http.ListenAndServe(address, nil))
}

func echoHandler(response http.ResponseWriter, request *http.Request) {
	connection, _ := upgrader.Upgrade(response, request, nil)
	defer connection.Close()

	reader := wsio.WebsocketReader{connection}
	writer := wsio.WebsocketWriter{connection}

	for {
		echo(&reader, &writer)
	}
}

func echo(reader io.Reader, writer io.Writer) {
	buffer := make([]byte, 1024)
	reader.Read(buffer)
	message := fmt.Sprintf("ECHO: %s", string(buffer))
	writer.Write([]byte(message))
}
