package websocket

type Hub struct {
	clients    map[*Client]struct{}
	register   chan *Client
	unregister chan *Client
	broadcast  chan Message
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]struct{}),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan Message),
	}
}

func (h *Hub) Run() {
	var uuid = 0
	for {
		select {
		case client := <-h.register:
			uuid++
			client.ID = string(uuid)
			h.clients[client] = struct{}{}
			go func() {
				client.send <- NewRegisterMessage(client.ID)
			}()
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				go func() {
					h.broadcast <- NewRemoveMessage(client.ID)
				}()
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					go func() {
						h.unregister <- client
					}()
				}
			}
		}
	}
}
