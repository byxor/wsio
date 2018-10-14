# wsio

[I/O](https://golang.org/pkg/io/) helpers to be used alongside [gorilla websockets](https://www.github.com/gorilla/websocket) in Go.

## Includes

* `WebsocketReader     // io.Reader implementation` 
* `WebsocketWriter     // io.Writer implementation`

## Why wsio?

These types allow you to write code that interacts with websockets, without depending on websockets.

By hiding the websocket's read/write calls behind interfaces, we can increase the portability of our code at a low cost.

<!--
## Before

## Af
-->
