package chat

type Hub struct {
	Connections map[*Connection]bool
	Register    chan *Connection
	Unregister  chan *Connection
	BroadCast   chan []byte
}

func (h *Hub) Start() {
	select {
	case c := <-h.Register:
		h.Connections[c] = true
	case c := <-h.Unregister:
		if _, ok := h.Connections[c]; ok {
			c.CloseChan <- []byte{}
			delete(h.Connections, c)
		}
	case m := <-h.BroadCast:
		for i := range h.Connections {
			i.SendChan <- m
		}
	}
}

var hub Hub = Hub{
	Connections: make(map[*Connection]bool),
	Register:    make(chan *Connection),
	Unregister:  make(chan *Connection),
	BroadCast:   make(chan []byte),
}
