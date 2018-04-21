package system

import (
	"cloudlike/component"
	"cloudlike/world"
)

// PlayerSystem .
func PlayerSystem(levels []*world.Level) {
	for _, level := range levels {
		for _, entity := range level.Entities {
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
							playerComponent.AddMessage("Wasn't given a direction to shoot!")
						} else {
							playerComponent.AddMessage("Shoot in the " + direction + " direction!")
						}
					}
				}
			}
		}
	}
}
