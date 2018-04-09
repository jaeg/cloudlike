package component

import "golang.org/x/net/websocket"

// WebSocketComponent - Handles websocket communications
type WebSocketComponent struct {
	ViewWidth, ViewHeight int
	Ws                    *websocket.Conn
	Command               map[string]interface{}
}

//GetType get the type
func (WebSocketComponent) GetType() string {
	return "WebSocketComponent"
}
