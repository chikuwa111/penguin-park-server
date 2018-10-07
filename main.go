package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chikuwa111/penguin-park/websocket"
)

func main() {
	hub := websocket.NewHub()
	go hub.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.Serve(hub, w, r)
	})

	port := os.Getenv("PORT")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
