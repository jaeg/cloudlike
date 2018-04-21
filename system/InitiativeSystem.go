package system

import (
	"cloudlike/component"
	"cloudlike/world"
)

// InitiativeSystem .
func InitiativeSystem(levels []*world.Level) {
	for _, level := range levels {
		for _, entity := range level.Entities {
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
}
