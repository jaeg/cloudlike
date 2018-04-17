package system

import (
	"cloudlike/entity"
	"cloudlike/world"
	"fmt"
)

// PlayerSystem .
func PlayerSystem(entities []*entity.Entity, levels []*world.Level) {
	for _, entity := range entities {
		if entity.HasComponent("PlayerComponent") {
			fmt.Println(entity.HasComponent("MyTurnComponent"))
			if entity.HasComponent("MyTurnComponent") {
				fmt.Println("It's a player's turn!")
			}
		}
	}
}
