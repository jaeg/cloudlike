package main

import (
	"cloudlike/component"
	"cloudlike/system"
)

func main() {
	helloWorldSystem := system.HelloWorldSystem{}
	helloWorldComponent := component.HelloWorldComponent{}

	helloWorldSystem.Update(&helloWorldComponent)
}
