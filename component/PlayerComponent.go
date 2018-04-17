package component

import "golang.org/x/net/websocket"

// PlayerComponent - Handles websocket communications
type PlayerComponent struct {
	ViewWidth, ViewHeight int
	Ws                    *websocket.Conn
	Command               map[string]interface{}
}

//GetType get the type
func (PlayerComponent) GetType() string {
	return "PlayerComponent"
}
