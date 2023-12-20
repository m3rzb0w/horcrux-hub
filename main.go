package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type application struct {
	conns map[*websocket.Conn]bool
}

func NewApplication() *application {
	return &application{
		conns: make(map[*websocket.Conn]bool),
	}
}

// func (s *Server) handleWS(ws *websocket.Conn) {
// 	fmt.Println("new incoming connection from client:", ws.RemoteAddr())
// 	s.conns[ws] = true

// 	s.readLoop(ws)
// }

// func (s *Server) readLoop(ws *websocket.Conn) {
// 	buf := make([]byte, 1024)
// 	for {
// 		n, err := ws.Read(buf)
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}
// 			fmt.Println("read error:", err)
// 			continue
// 		}
// 		msg := buf[:n]

// 		s.broadCast(msg)
// 	}
// }

// func (s *Server) broadCast(b []byte) {
// 	for ws := range s.conns {
// 		go func(ws *websocket.Conn) {
// 			if _, err := ws.Write(b); err != nil {
// 				fmt.Println("write error:", err)
// 			}
// 		}(ws)
// 	}
// }

const port = 8080

func main() {

	var app = NewApplication()
	log.Println("Starting application on port", port)

	// start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}

}
