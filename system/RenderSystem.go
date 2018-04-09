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
func RenderSystem(entities *[]*entity.Entity, levels *[]world.Level) {
	fmt.Println(entities)
	for _, entity := range *entities {
		fmt.Println(entity)
		if entity.HasComponent("WebSocketComponent") {
			// Look up level by id
			// Construct view of
			fmt.Println("Rendered")
			wsc := (*entity.GetComponent("WebSocketComponent")).(*component.WebSocketComponent)
			viewWidth := wsc.ViewWidth
			viewHeight := wsc.ViewHeight
			positionComponent := (*entity.GetComponent("PositionComponent")).(*component.PositionComponent)
			onLevel := (*levels)[positionComponent.Level]

			view := onLevel.GetView(positionComponent.X, positionComponent.Y, viewWidth, viewHeight, false)
			viewJSONBytes, _ := json.Marshal(view)
			n := len(viewJSONBytes)
			viewJSON := string(viewJSONBytes[:n])
			if err := websocket.Message.Send(wsc.Ws, "view:"+viewJSON); err != nil {
				//fmt.Printf("Can't send to player %v\n", i)
			}
		}
	}
}
