package system

import (
	"cloudlike/component"
	"cloudlike/world"
)

var statusConditions = []string{"Poisoned"}

// StatusConditionSystem .
func StatusConditionSystem(levels []*world.Level) {
	for _, level := range levels {
		for _, entity := range level.Entities {
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
}
