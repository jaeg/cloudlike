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

var entities []entity.Entity
var levels []world.Level

func main() {
	entities = []entity.Entity{}
	levels = []world.Level{}
	levels = append(levels, *world.NewOverworldSection(16, 16))

	ticker := time.NewTicker(time.Second / 4)
	go func() {
		for _ = range ticker.C {
			system.RenderSystem(entities, &levels)
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
	webSocketComponent := &component.WebSocketComponent{Ws: ws}
	newPlayerEntity.AddComponent(webSocketComponent)
	positionComponent := &component.PositionComponent{X: 0, Y: 0, Level: 0}
	newPlayerEntity.AddComponent(positionComponent)
	entities = append(entities, newPlayerEntity)

	fmt.Println("Entities handle sock:", entities)
	system.WebSocketSystem(newPlayerEntity, ws)
}
