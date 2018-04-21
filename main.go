package main

import (
	"cloudlike/component"
	"cloudlike/entity"
	"cloudlike/system"
	"cloudlike/world"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

const port = "1234"

var levels []*world.Level

func main() {
	levels = []*world.Level{}
	levels = append(levels, world.NewOverworldSection(16, 16))

	ticker := time.NewTicker(time.Second / 4)
	go func() {
		for _ = range ticker.C {
			system.InitiativeSystem(levels)
			system.PlayerSystem(levels)
			system.RenderSystem(levels)
			system.StatusConditionSystem(levels)
			levels = system.CleanUpSystem(levels)
		}

	}()

	//Setup the websocket handler
	fmt.Println("Setting up websocket handler")
	http.Handle("/", websocket.Handler(HandleSocket))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	} else {
		fmt.Println("Listening on " + port)
	}
}

func HandleSocket(ws *websocket.Conn) {
	// Create player since this is a new connection
	fmt.Println("New client connected.")
	newPlayerEntity := entity.Entity{}
	playerComponent := &component.PlayerComponent{Ws: ws}
	newPlayerEntity.AddComponent(playerComponent)

	poison := &component.PoisonedComponent{Duration: 5}
	newPlayerEntity.AddComponent(poison)
	initiativeComponent := &component.InitiativeComponent{DefaultValue: 1, Ticks: 1}
	newPlayerEntity.AddComponent(initiativeComponent)
	positionComponent := &component.PositionComponent{X: 0, Y: 0, Level: 0}
	newPlayerEntity.AddComponent(positionComponent)
	newPlayerEntity.AddComponent(&component.AppearanceComponent{SpriteIndex: 0, Resource: "npc"})
	newPlayerEntity.AddComponent(&component.DirectionComponent{Direction: 0})
	//entities = append(entities, &newPlayerEntity)
	levels[0].AddEntity(&newPlayerEntity)

	system.WebSocketSystem(newPlayerEntity, ws)
}
