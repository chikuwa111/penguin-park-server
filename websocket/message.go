package websocket

type Message struct {
	Type      string  `json:"type"`
	ID        string  `json:"id"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Scale     float64 `json:"scale"`
	Direction string  `json:"direction"`
}

func NewRegisterMessage(id string) Message {
	return Message{
		Type: "REGISTER",
		ID:   id,
	}
}

func NewRemoveMessage(id string) Message {
	return Message{
		Type: "REMOVE",
		ID:   id,
	}
}
