# wsio

[I/O](https://golang.org/pkg/io/) helpers to be used alongside [gorilla websockets](https://www.github.com/gorilla/websocket) in Go.

[![Documentation](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/byxor/wsio)

## Includes

* `WebsocketReader` (implements [io.Reader](https://golang.org/pkg/io/#Reader))
* `WebsocketWriter` (implements [io.Writer](https://golang.org/pkg/io/#Writer))

## Why wsio?

* It allows you to write code that interacts with websockets, without depending on websockets.
* It increases the testability of your code.
* It reduces the cost of switching networking implementations.

## Before

```go
package main

import (
	"fmt"
	"github.com/gorilla/websocket"
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

	for {
		echo(connection)
	}
}

func echo(connection *websocket.Conn) {
	messageType, data, _ := connection.ReadMessage()
	message := fmt.Sprintf("ECHO: %s", string(data))
	connection.WriteMessage(messageType, []byte(message))
}
```

## After

```go
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
	message := fmt.Sprintf("ECHO: %s", string(buffer))
	writer.Write([]byte(message))
}
```
