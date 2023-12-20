package main

import (
	"fmt"
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
