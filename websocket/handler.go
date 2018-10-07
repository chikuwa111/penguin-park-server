package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	allowOrigins = []string{"https://penguins-park.firebaseapp.com"}
	upgrader     = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			origins, ok := r.Header["Origin"]
			if !ok {
				return false
			}
			origin := origins[0]
			for _, allowOrigin := range allowOrigins {
				if origin == allowOrigin {
					return true
				}
			}
			return false
		},
	}
)

func Serve(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := NewClient(hub, conn)
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}
