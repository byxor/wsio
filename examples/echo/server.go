package main

import (
	"fmt"
	"github.com/byxor/wsio"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", echoHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func echoHandler(response http.ResponseWriter, request *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

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
	input := string(buffer)
	output := fmt.Sprintf("ECHO: %s", input)
	writer.Write([]byte(output))
}
