package system

import (
	"cloudlike/component"
	"cloudlike/entity"
	"encoding/json"
	"fmt"

	"golang.org/x/net/websocket"
)

// WebSocketSystem .
type WebSocketSystem struct {
	Entities *[]entity.Entity
}

//HandleSocket Handles player input from the client and puts data into a component to be handled during the turn loop
func (wSS *WebSocketSystem) HandleSocket(ws *websocket.Conn) {
	// Create player since this is a new connection
	fmt.Println("New client connected.")
	newPlayerEntity := &entity.Entity{}
	webSocketComponent := &component.WebSocketComponent{Ws: ws}
	newPlayerEntity.AddComponent(webSocketComponent)
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

		webSocketComponent.Command = command
		fmt.Println(*(newPlayerEntity.Components[0]))
	}
}
