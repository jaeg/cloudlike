package system

import (
	"cloudlike/component"
	"cloudlike/world"
	"math/rand"
)

func getRandom(low int, high int) int {
	return (rand.Intn((high - low))) + low
}

// PlayerSystem .
func AISystem(levels []*world.Level) {
	for _, level := range levels {
		//fmt.Println(t, len(level.Entities))
		for _, entity := range level.Entities {

			if entity.HasComponent("WanderAIComponent") {
				if entity.HasComponent("MyTurnComponent") {
					pc := entity.GetComponent("PositionComponent").(*component.PositionComponent)
					dc := entity.GetComponent("DirectionComponent").(*component.DirectionComponent)

					deltaX := getRandom(-1, 2)
					deltaY := 0
					if deltaX == 0 {
						deltaY = getRandom(-1, 2)
					}
					if level.GetSolidEntityAt(pc.X+deltaX, pc.Y+deltaY) == nil {
						tile := level.GetTileAt(pc.X+deltaX, pc.Y+deltaY)
						if tile == nil {
							tile = level.GetTileAt(pc.X, pc.Y)
							if tile.VertTo.Level >= 0 {
								//Switch level
								level.RemoveEntity(entity)
								levels[tile.VertTo.Level].AddEntity(entity)
								pc.Level = tile.VertTo.Level
								pc.X = tile.VertTo.X
								pc.Y = tile.VertTo.Y
							}
						} else if tile.Solid == false && tile.Water == false {
							pc.X += deltaX
							pc.Y += deltaY
						}
					}
					if deltaY > 0 {
						dc.Direction = 1
					}
					if deltaY < 0 {
						dc.Direction = 2
					}
					if deltaX < 0 {
						dc.Direction = 3
					}
					if deltaX > 0 {
						dc.Direction = 0
					}
				}
			}
		}
	}
}
