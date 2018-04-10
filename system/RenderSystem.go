package system

import (
	"cloudlike/component"
	"cloudlike/entity"
	"cloudlike/world"
	"encoding/json"
	"fmt"

	"golang.org/x/net/websocket"
)

// RenderSystem .
func RenderSystem(entities []entity.Entity, levels *[]world.Level) {
	for _, entity := range entities {
		if entity.HasComponent("WebSocketComponent") {
			// Look up level by id
			// Construct view of
			wsc := entity.GetComponent("WebSocketComponent").(*component.WebSocketComponent)
			viewWidth := wsc.ViewWidth
			viewHeight := wsc.ViewHeight
			positionComponent := entity.GetComponent("PositionComponent").(*component.PositionComponent)
			onLevel := (*levels)[positionComponent.Level]

			view := onLevel.GetView(positionComponent.X, positionComponent.Y, viewWidth, viewHeight, false)
			viewJSONBytes, _ := json.Marshal(view)
			n := len(viewJSONBytes)
			viewJSON := string(viewJSONBytes[:n])
			if wsc.Ws != nil {
				if err := websocket.Message.Send(wsc.Ws, "view:"+viewJSON); err != nil {
					fmt.Printf("Can't send to player %v\n", entity)
				}
			}

			//fmt.Println(viewJSON)
		}
	}
}
