package system

import (
	"cloudlike/component"
	"fmt"
)

// HelloWorldSystem .
type HelloWorldSystem struct {
}

// Update .
func (HelloWorldSystem) Update(a *component.HelloWorldComponent) {
	fmt.Println("hello world")
}
