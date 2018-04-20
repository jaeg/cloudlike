package component

import "golang.org/x/net/websocket"

// PlayerComponent - Handles websocket communications
type PlayerComponent struct {
	ViewWidth, ViewHeight int
	Ws                    *websocket.Conn
	Commands              []string
	MessageLog            []string
}

//GetType get the type
func (PlayerComponent) GetType() string {
	return "PlayerComponent"
}

func (pc *PlayerComponent) PushCommand(x string) {
	pc.Commands = append(pc.Commands, x)
}

func (pc *PlayerComponent) AddMessage(x string) {
	pc.MessageLog = append(pc.MessageLog, x)
}

func (pc *PlayerComponent) PopCommand() string {
	x := ""
	if len(pc.Commands) > 0 {
		x, pc.Commands = pc.Commands[0], pc.Commands[1:]
	}

	return x
}
