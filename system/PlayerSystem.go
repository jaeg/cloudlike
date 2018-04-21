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
						if level.GetSolidEntityAt(pc.X, pc.Y-1) == nil {
							pc.Y--
							dc.Direction = 2
						} else {
							playerComponent.AddMessage("Something blocks you in that direction!")
						}
					case "S":
						if level.GetSolidEntityAt(pc.X, pc.Y+1) == nil {
							pc.Y++
							dc.Direction = 1
						} else {
							playerComponent.AddMessage("Something blocks you in that direction!")
						}
					case "A":
						if level.GetSolidEntityAt(pc.X-1, pc.Y) == nil {
							pc.X--
							dc.Direction = 3
						} else {
							playerComponent.AddMessage("Something blocks you in that direction!")
						}
					case "D":
						if level.GetSolidEntityAt(pc.X+1, pc.Y) == nil {
							pc.X++
							dc.Direction = 0
						} else {
							playerComponent.AddMessage("Something blocks you in that direction!")
						}
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
