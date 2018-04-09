package entity

import (
	"cloudlike/component"
	"fmt"
)

// Entity .
type Entity struct {
	Components []*component.Component
}

func (entity *Entity) AddComponent(component component.Component) {
	entity.Components = append(entity.Components, &component)
	fmt.Println(entity)
}
