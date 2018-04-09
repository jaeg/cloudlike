package main

import (
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

func main() {
	entities := []entity.Entity{}
	levels := []world.Level{}
	levels = append(levels, *world.NewOverworldSection(16, 16))

	ticker := time.NewTicker(time.Second / 4)
	go func() {
		for _ = range ticker.C {
			system.RenderSystem(entities, &levels)
		}

	}()

	//Setup the websocket handler
	webSocketSystem := &system.WebSocketSystem{Entities: &entities}

	fmt.Println("Setting up websocket handler")
	http.Handle("/", websocket.Handler(webSocketSystem.HandleSocket))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	} else {
		fmt.Println("Listening on " + port)
	}
}
