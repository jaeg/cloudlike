package system

import (
	"cloudlike/component"
	"cloudlike/entity"
	"cloudlike/world"
	"fmt"
)

// PlayerSystem .
func PlayerSystem(entities []*entity.Entity, levels []*world.Level) {
	for _, entity := range entities {
		if entity.HasComponent("PlayerComponent") {
			if entity.HasComponent("MyTurnComponent") {
				pc := entity.GetComponent("PositionComponent").(*component.PositionComponent)
				dc := entity.GetComponent("DirectionComponent").(*component.DirectionComponent)
				playerComponent := entity.GetComponent("PlayerComponent").(*component.PlayerComponent)
				command := playerComponent.PopCommand()
				switch command {
				case "W":
					pc.Y--
					dc.Direction = 2
				case "S":
					pc.Y++
					dc.Direction = 1
				case "A":
					pc.X--
					dc.Direction = 3
				case "D":
					pc.X++
					dc.Direction = 0
				case "F":
					direction := playerComponent.PopCommand()
					if direction == "" {
						fmt.Println("Wasn't given a direction to shoot!")
					} else {
						fmt.Println("Shoot in the " + direction + " direction!")
					}
				}
			}
		}
	}
}
