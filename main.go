package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type application struct {
	conns     map[*websocket.Conn]bool
	broadcast chan Message
}

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func NewApplication() *application {
	return &application{
		conns:     make(map[*websocket.Conn]bool),
		broadcast: make(chan Message),
	}
}

const port = 8080

func main() {

	var app = NewApplication()
	log.Println("Starting application on port", port)
	go app.HandleMessages()
	// start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}

}
