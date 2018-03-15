package system

import (
	"cloudlike/component"
	"fmt"
)

type HelloWorldSystem struct {
}

func (HelloWorldSystem) Update(a *component.HelloWorldComponent) {
	fmt.Println("hello world")
}
