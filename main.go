package main

import (
	"cloudlike/entity"
	"cloudlike/system"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

const port = "1234"

func main() {
	entities := []entity.Entity{}
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
