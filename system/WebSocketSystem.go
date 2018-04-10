package system

import (
	"cloudlike/component"
	"cloudlike/entity"
	"encoding/json"
	"fmt"

	"golang.org/x/net/websocket"
)

//HandleSocket Handles player input from the client and puts data into a component to be handled during the turn loop
func WebSocketSystem(entity entity.Entity, ws *websocket.Conn) {
	var err error
	if !entity.HasComponent("WebSocketComponent") {
		return
	}

	webSocketComponent := entity.GetComponent("WebSocketComponent").(*component.WebSocketComponent)
	// Message loop
	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Printf("Socket Error: %v \n", err)
			webSocketComponent.Ws = nil
			break
		}

		var command map[string]interface{}
		if err = json.Unmarshal([]byte(reply), &command); err != nil {
			fmt.Println("Issue unmarshaling " + reply)
			continue
		}
		fmt.Println(command)

		switch command["type"] {
		case "viewSize":
			webSocketComponent.ViewWidth = int(command["width"].(float64))
			webSocketComponent.ViewHeight = int(command["height"].(float64))
		default:
			webSocketComponent.Command = command
		}
	}
}
