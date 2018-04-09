package system

import (
	"cloudlike/component"
	"cloudlike/entity"
	"encoding/json"
	"fmt"

	"golang.org/x/net/websocket"
)

// WebSocketSystem This system handles websockets for the game.
type WebSocketSystem struct {
	Entities *[]entity.Entity
}

//HandleSocket Handles player input from the client and puts data into a component to be handled during the turn loop
func (wSS *WebSocketSystem) HandleSocket(ws *websocket.Conn) {
	// Create player since this is a new connection
	fmt.Println("New client connected.")
	newPlayerEntity := entity.Entity{}
	webSocketComponent := &component.WebSocketComponent{Ws: ws}
	newPlayerEntity.AddComponent(webSocketComponent)
	positionComponent := &component.PositionComponent{X: 0, Y: 0, Level: 0}
	newPlayerEntity.AddComponent(positionComponent)
	*wSS.Entities = append(*wSS.Entities, newPlayerEntity)
	fmt.Println(newPlayerEntity)

	var err error

	// Message loop
	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Printf("Socket Error: %v \n", err)
			break
		}

		var command map[string]interface{}
		if err = json.Unmarshal([]byte(reply), &command); err != nil {
			fmt.Println("Issue unmarshaling " + reply)
			continue
		}

		switch command["type"] {
		case "viewSize":
			webSocketComponent.ViewWidth = int(command["width"].(float64))
			webSocketComponent.ViewHeight = int(command["height"].(float64))
		default:
			webSocketComponent.Command = command
		}

		fmt.Println(newPlayerEntity.Components[0])
	}
}
