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
				fmt.Println("It's a player's turn!")
				pc := entity.GetComponent("PositionComponent").(*component.PositionComponent)
				playerComponent := entity.GetComponent("PlayerComponent").(*component.PlayerComponent)
				command := playerComponent.PopCommand()
				switch command {
				case "W":
					pc.Y--
				case "S":
					pc.Y++
				case "A":
					pc.X--
				case "D":
					pc.X++
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
