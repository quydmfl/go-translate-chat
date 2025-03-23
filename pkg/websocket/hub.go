package websocket

import (
	"log"
	"sync"
)

type Hub struct {
	Clients    map[string]*Client            // List user online
	Broadcast  chan Message                  // Broadcast channel
	Register   chan *Client                  // New client register
	Unregister chan *Client                  // Remove client leave
	Rooms      map[string]map[string]*Client // list room chat with key is room name
	sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[string]*Client),
		Broadcast:  make(chan Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Rooms:      make(map[string]map[string]*Client),
	}
}

func (h *Hub) Run() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Hub recovered from panic: %v", r)
			go h.Run() // Restart the loop
		}
	}()

	for {
		select {
		case client := <-h.Register:
			h.Clients[client.Username] = client

		case client := <-h.Unregister:
			if _, ok := h.Clients[client.Username]; ok {
				delete(h.Clients, client.Username)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			h.HandleMessage(message)
		}
	}
}

// HandleMessage //
func (h *Hub) HandleMessage(message Message) {
	switch message.Type {
	case "private":
		// private message
		if receiver, ok := h.Clients[message.Target]; ok {
			receiver.Send <- message
		}

	case "group":
		// send to group
		h.SendToGroup(message.Target, message)

	case "join_group":
		// add user to group
		h.AddClientToGroup(message.Target, message.Sender)
	}
}

func (h *Hub) AddClientToGroup(groupName, username string) {
	h.Lock()
	defer h.Unlock()

	if _, ok := h.Rooms[groupName]; !ok {
		h.Rooms[groupName] = make(map[string]*Client)
	}

	if client, exists := h.Clients[username]; exists {
		h.Rooms[groupName][username] = client
	}
}

func (h *Hub) SendToGroup(groupName string, message Message) {
	h.Lock()
	defer h.Unlock()

	if group, ok := h.Rooms[groupName]; ok {
		for _, client := range group {
			client.Send <- message
		}
	}
}
