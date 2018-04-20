package system

import (
	"cloudlike/component"
	"cloudlike/entity"
	"cloudlike/world"
)

var statusConditions = []string{"Poisoned"}

// StatusConditionSystem .
func StatusConditionSystem(entities []*entity.Entity, levels []*world.Level) {
	for _, entity := range entities {
		for _, statusCondition := range statusConditions {
			if entity.HasComponent(statusCondition + "Component") {
				pc := entity.GetComponent(statusCondition + "Component").(component.DecayingComponent)

				if pc.Decay() {
					entity.RemoveComponent(statusCondition + "Component")
					if entity.HasComponent("PlayerComponent") {
						playerComponent := entity.GetComponent("PlayerComponent").(*component.PlayerComponent)
						playerComponent.AddMessage(statusCondition + " has cleared!")
					}
				}
			}
		}
	}
}
