package system

import (
	"cloudlike/world"
	"fmt"
)

// CleanUpSystem .
func CleanUpSystem(levels []*world.Level) []*world.Level {
	for _, level := range levels {
		for i, entity := range level.Entities {
			if entity.HasComponent("MyTurnComponent") {
				entity.RemoveComponent("MyTurnComponent")
			}

			if entity.HasComponent("DeadComponent") {
				level.Entities = append(level.Entities[:i], level.Entities[i+1:]...)
				fmt.Println("Killed")
			}

		}
	}

	return levels
}
