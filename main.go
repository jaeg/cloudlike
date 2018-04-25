package main

import (
	"cloudlike/component"
	"cloudlike/entity"
	"cloudlike/system"
	"cloudlike/world"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

const port = "1234"

var levels []*world.Level

func main() {
	levels = []*world.Level{}
	levels = append(levels, world.NewOverworldSection(100, 100))
	levels = append(levels, world.NewOverworldSection(10, 10))

	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			if rand.Intn(100) == 0 {
				entity := entity.Entity{}
				entity.AddComponent(&component.WanderAIComponent{})
				entity.AddComponent(&component.InitiativeComponent{DefaultValue: 4, Ticks: 1})
				entity.AddComponent(&component.PositionComponent{X: x, Y: y, Level: 0})
				entity.AddComponent(&component.AppearanceComponent{SpriteIndex: 4, Resource: "npc"})
				entity.AddComponent(&component.DirectionComponent{Direction: 0})
				entity.AddComponent(&component.SolidComponent{})
				//entities = append(entities, &newPlayerEntity)
				levels[0].AddEntity(&entity)
			}
		}
	}

	ticker := time.NewTicker(time.Second / 4)
	go func() {
		for _ = range ticker.C {
			system.InitiativeSystem(levels)
			system.PlayerSystem(levels)
			system.AISystem(levels)
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
	newPlayerEntity.AddComponent(&component.SolidComponent{})
	//entities = append(entities, &newPlayerEntity)
	levels[0].AddEntity(&newPlayerEntity)

	system.WebSocketSystem(newPlayerEntity, ws)
}
