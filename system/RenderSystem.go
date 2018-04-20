package system

import (
	"cloudlike/component"
	"cloudlike/entity"
	"cloudlike/world"
	"encoding/json"
	"fmt"

	"golang.org/x/net/websocket"
)

type entityView struct {
	X, Y        int
	SpriteIndex int
	Resource    string
	Direction   int
}

// RenderSystem .
func RenderSystem(entities []*entity.Entity, levels []*world.Level) {

	var seeableEntities []entityView
	for _, entity := range entities {
		if entity.HasComponent("AppearanceComponent") {
			ac := entity.GetComponent("AppearanceComponent").(*component.AppearanceComponent)
			pc := entity.GetComponent("PositionComponent").(*component.PositionComponent)
			direction := -1
			if entity.HasComponent("DirectionComponent") {
				dc := entity.GetComponent("DirectionComponent").(*component.DirectionComponent)
				direction = dc.Direction
			}
			ev := entityView{X: pc.X, Y: pc.Y, SpriteIndex: ac.SpriteIndex, Resource: ac.Resource, Direction: direction}
			seeableEntities = append(seeableEntities, ev)
		}
	}

	viewJSONBytes, _ := json.Marshal(seeableEntities)
	n := len(viewJSONBytes)
	entitiesJSON := string(viewJSONBytes[:n])

	for _, entity := range entities {
		if entity.HasComponent("PlayerComponent") {
			// Look up level by id
			// Construct view of
			wsc := entity.GetComponent("PlayerComponent").(*component.PlayerComponent)
			viewWidth := wsc.ViewWidth
			viewHeight := wsc.ViewHeight
			positionComponent := entity.GetComponent("PositionComponent").(*component.PositionComponent)
			onLevel := levels[positionComponent.Level]

			view := onLevel.GetView(positionComponent.X, positionComponent.Y, viewWidth, viewHeight, false)
			viewJSONBytes, _ := json.Marshal(view)
			n := len(viewJSONBytes)
			viewJSON := string(viewJSONBytes[:n])

			playerJSONBytes, _ := json.Marshal(positionComponent)
			n = len(playerJSONBytes)
			playerJSON := string(playerJSONBytes[:n])

			playerCommandQueueBytes, _ := json.Marshal(wsc.Commands)
			n = len(playerCommandQueueBytes)
			playerCommandQueueJSON := string(playerCommandQueueBytes[:n])

			playerMessageLogBytes, _ := json.Marshal(wsc.MessageLog)
			n = len(playerMessageLogBytes)
			playerMessageLogJSON := string(playerMessageLogBytes[:n])

			if wsc.Ws != nil {
				if err := websocket.Message.Send(wsc.Ws, "view:"+viewJSON); err != nil {
					fmt.Printf("Can't send view to player %v\n", entity)
				}

				if err := websocket.Message.Send(wsc.Ws, "entities:"+entitiesJSON); err != nil {
					fmt.Printf("Can't send entities to player %v\n", entity)
				}

				if err := websocket.Message.Send(wsc.Ws, "player:"+playerJSON); err != nil {
					fmt.Printf("Can't send player to player %v\n", entity)
				}

				if err := websocket.Message.Send(wsc.Ws, "commandQueue:"+playerCommandQueueJSON); err != nil {
					fmt.Printf("Can't send command queue to player %v\n", entity)
				}

				if err := websocket.Message.Send(wsc.Ws, "messageLog:"+playerMessageLogJSON); err != nil {
					fmt.Printf("Can't send command queue to player %v\n", entity)
				}
			}
		}
	}
}
