package system

import (
	"cloudlike/entity"
	"cloudlike/world"
	"fmt"
)

// CleanUpSystem .
func CleanUpSystem(entities []*entity.Entity, levels []*world.Level) ([]*entity.Entity, []*world.Level) {
	for i, entity := range entities {
		if entity.HasComponent("MyTurnComponent") {
			entity.RemoveComponent("MyTurnComponent")
		}

		if entity.HasComponent("DeadComponent") {
			entities = append(entities[:i], entities[i+1:]...)
			fmt.Println("Killed")
		}

	}

	return entities, levels
}
