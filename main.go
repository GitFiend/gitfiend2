package main

import (
	"gitfiend2/core/server"
)

//go:generate go run core/parser/genand/main.go
//go:generate gofmt -w core/parser

func main() {
	server.StartServer()
}
