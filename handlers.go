package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	frontend := getFrontendAssets()
	fs := http.FileServer(http.FS(frontend))
	fs.ServeHTTP(w, r)
}

func (app *application) Ping(w http.ResponseWriter, r *http.Request) {
	ping := fmt.Sprintf("Ping : %v", time.Now())
	w.Write([]byte(ping))
}

func (a *application) handleWSOrderbook(ws *websocket.Conn) {
	fmt.Println("new incomming connection from client to countdown feed", ws.RemoteAddr())
	for {
		deadline := time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)
		diff := deadline.Sub(time.Now())
		payload := fmt.Sprintf("%s\n", diff.Round(time.Second))
		if diff < 0 {
			payload = fmt.Sprintf("Time to give back the ps4 ! Additional Keeping time : %s\n", diff.Abs().Round(time.Second))
		}
		ws.Write([]byte(payload))
		time.Sleep(time.Second * 1)
	}
}

func (a *application) handleWS(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client:", ws.RemoteAddr())
	a.conns[ws] = true

	a.readLoop(ws)
}

func (a *application) readLoop(ws *websocket.Conn) {
	defer ws.Close()
	buf := make([]byte, 1024)
	for {
		var msg Message
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error:", err)
			continue
		}

		err = json.Unmarshal(buf[:n], &msg)
		if err != nil {
			fmt.Println(err)
			delete(a.conns, ws)
			return
		}
		a.broadcast <- msg
	}
}

func (a *application) HandleMessages() {
	for {
		msg := <-a.broadcast
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}
		for ws := range a.conns {
			func(ws *websocket.Conn) {
				if _, err := ws.Write(msgBytes); err != nil {
					fmt.Println("write error:", err)
					ws.Close()
					delete(a.conns, ws)
				}
			}(ws)
		}
	}
}
