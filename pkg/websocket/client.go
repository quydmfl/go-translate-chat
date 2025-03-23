package websocket

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	Username string
	Conn     *websocket.Conn
	Send     chan Message
	Hub      *Hub
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		var message Message

		_, rawMsg, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		fmt.Println("Raw JSON message:", string(rawMsg))

		err = json.Unmarshal(rawMsg, &message)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			continue
		}

		if message.Type == "" || message.Sender == "" {
			fmt.Println("Invalid message format:", message)
			continue
		}

		c.Hub.Broadcast <- message
	}
}

func (c *Client) WritePump() {
	defer c.Conn.Close()
	for message := range c.Send {
		c.Conn.WriteJSON(message)
	}
}
