package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chikuwa111/penguin-park-server/websocket"
)

func main() {
	hub := websocket.NewHub()
	go hub.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.Serve(hub, w, r)
	})

	port := os.Getenv("PORT")
	cert := os.Getenv("CERT")
	key := os.Getenv("KEY")
	if err := http.ListenAndServeTLS(":"+port, cert, key, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
