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

func (entity *Entity) HasComponent(name string) bool {
	for _, component := range entity.Components {
		if (*component).GetType() == name {
			return true
		}
	}
	return false
}

func (entity *Entity) GetComponent(name string) *component.Component {
	for _, component := range entity.Components {
		if (*component).GetType() == name {
			return component
		}
	}
	return nil
}
