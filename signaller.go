package main

// Channel allows clients to subscribe to different topics
type Channel struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func newChannel() *Channel {
	return &Channel{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (s *Channel) run() {
	for {
		select {
		case client := <-s.register:
			s.clients[client] = true

		case client := <-s.unregister:
			if _, ok := s.clients[client]; ok {
				delete(s.clients, client)
				close(client.send)
			}

		case message := <-s.broadcast:
			for client := range s.clients {
				select {
				case client.send <- message:
				default: // If we can't send, presume the client has DC'd
					close(client.send)
					delete(s.clients, client)
				}
			}
		}
	}
}
