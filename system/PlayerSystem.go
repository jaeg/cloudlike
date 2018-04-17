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
				pc.X++
			}
		}
	}
}
