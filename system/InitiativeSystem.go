package system

import (
	"cloudlike/component"
	"cloudlike/entity"
	"cloudlike/world"
)

// InitiativeSystem .
func InitiativeSystem(entities []*entity.Entity, levels []*world.Level) {
	for _, entity := range entities {
		if entity.HasComponent("InitiativeComponent") {
			ic := entity.GetComponent("InitiativeComponent").(*component.InitiativeComponent)
			ic.Ticks--

			if ic.Ticks <= 0 {
				ic.Ticks = ic.DefaultValue
				if ic.OverrideValue > 0 {
					ic.Ticks = ic.OverrideValue
				}

				if entity.HasComponent("MyTurnComponent") == false {
					mTC := &component.MyTurnComponent{}
					entity.AddComponent(mTC)
				}
			}
		}
	}
}
