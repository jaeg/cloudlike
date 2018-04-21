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
						if level.GetEntityAt(pc.X, pc.Y-1) == nil {
							pc.Y--
							dc.Direction = 2
						}
					case "S":
						if level.GetEntityAt(pc.X, pc.Y+1) == nil {
							pc.Y++
							dc.Direction = 1
						}
					case "A":
						if level.GetEntityAt(pc.X-1, pc.Y) == nil {
							pc.X--
							dc.Direction = 3
						}
					case "D":
						if level.GetEntityAt(pc.X+1, pc.Y) == nil {
							pc.X++
							dc.Direction = 0
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
