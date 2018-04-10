package entity

import (
	"cloudlike/component"
	"fmt"
)

// Entity .
type Entity struct {
	Components []component.Component
}

func (entity *Entity) AddComponent(component component.Component) {
	entity.Components = append(entity.Components, component)
	fmt.Println(entity)
}

func (entity *Entity) HasComponent(name string) bool {
	for _, component := range entity.Components {
		if (component).GetType() == name {
			return true
		}
	}
	return false
}

func (entity *Entity) GetComponent(name string) component.Component {
	for _, component := range entity.Components {
		if component.GetType() == name {
			return component
		}
	}
	return nil
}

func (entity *Entity) RemoveComponent(name string) {
	for i, component := range entity.Components {
		if component.GetType() == name {
			copy(entity.Components[i:], entity.Components[i+1:])
			entity.Components[len(entity.Components)-1] = nil // or the zero value of T
			entity.Components = entity.Components[:len(entity.Components)-1]
		}
	}

}
