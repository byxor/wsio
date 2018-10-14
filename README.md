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

<!--
## Before

```go

```

## After

```go

```
-->
