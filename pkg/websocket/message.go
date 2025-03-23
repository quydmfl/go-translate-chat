package websocket

type Message struct {
	Type    string `json:"type"` // "private", "group", "broadcast"
	Sender  string `json:"sender"`
	Target  string `json:"target"` // username or group name
	Message string `json:"message,omitempty"`
}
