package system

import (
	"cloudlike/entity"
	"cloudlike/world"
)

// CleanUpSystem .
func CleanUpSystem(entities []*entity.Entity, levels []*world.Level) {
	for _, entity := range entities {

		if entity.HasComponent("MyTurnComponent") {
			entity.RemoveComponent("MyTurnComponent")
		}

	}
}
