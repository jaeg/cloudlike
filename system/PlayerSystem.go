package system

import (
	"cloudlike/component"
	"cloudlike/world"
)

// PlayerSystem .
func PlayerSystem(levels []*world.Level) {
	for _, level := range levels {
		//fmt.Println(t, len(level.Entities))
		for _, entity := range level.Entities {

			if entity.HasComponent("PlayerComponent") {
				if entity.HasComponent("MyTurnComponent") {
					pc := entity.GetComponent("PositionComponent").(*component.PositionComponent)
					dc := entity.GetComponent("DirectionComponent").(*component.DirectionComponent)
					playerComponent := entity.GetComponent("PlayerComponent").(*component.PlayerComponent)
					command := playerComponent.PopCommand()
					switch command {
					case "W":
						if level.GetSolidEntityAt(pc.X, pc.Y-1) == nil {
							tile := level.GetTileAt(pc.X, pc.Y-1)
							if tile == nil {
								tile = level.GetTileAt(pc.X, pc.Y)
								if tile.VertTo.Level == -1 {
									playerComponent.AddMessage("You've hit the edge of the world!")
								} else {
									//Switch level
									level.RemoveEntity(entity)
									levels[tile.VertTo.Level].AddEntity(entity)
									pc.Level = tile.VertTo.Level
									pc.X = tile.VertTo.X
									pc.Y = tile.VertTo.Y
								}
							} else if tile.Solid == true || tile.Water == true {
								playerComponent.AddMessage("You can't walk that way!")
							} else {
								pc.Y--

							}
						} else {
							playerComponent.AddMessage("Something blocks you in that direction!")
						}
						dc.Direction = 2
					case "S":
						if level.GetSolidEntityAt(pc.X, pc.Y+1) == nil {
							tile := level.GetTileAt(pc.X, pc.Y+1)
							if tile == nil {
								tile = level.GetTileAt(pc.X, pc.Y)
								if tile.VertTo.Level == -1 {
									playerComponent.AddMessage("You've hit the edge of the world!")
								} else {
									level.RemoveEntity(entity)
									levels[tile.VertTo.Level].AddEntity(entity)
									pc.Level = tile.VertTo.Level
									pc.X = tile.VertTo.X
									pc.Y = tile.VertTo.Y
								}
							} else if tile.Solid == true || tile.Water == true {
								playerComponent.AddMessage("You can't walk that way!")
							} else {
								pc.Y++

							}
						} else {
							playerComponent.AddMessage("Something blocks you in that direction!")
						}
						dc.Direction = 1
					case "A":
						if level.GetSolidEntityAt(pc.X-1, pc.Y) == nil {
							tile := level.GetTileAt(pc.X-1, pc.Y)
							if tile == nil {
								tile = level.GetTileAt(pc.X, pc.Y)
								if tile.HorzTo.Level == -1 {
									playerComponent.AddMessage("You've hit the edge of the world!")
								} else {
									level.RemoveEntity(entity)
									levels[tile.HorzTo.Level].AddEntity(entity)
									pc.Level = tile.HorzTo.Level
									pc.X = tile.HorzTo.X
									pc.Y = tile.HorzTo.Y
								}
							} else if tile.Solid == true || tile.Water == true {
								playerComponent.AddMessage("You can't walk that way!")
							} else {
								pc.X--

							}
						} else {
							playerComponent.AddMessage("Something blocks you in that direction!")
						}
						dc.Direction = 3
					case "D":
						if level.GetSolidEntityAt(pc.X+1, pc.Y) == nil {
							tile := level.GetTileAt(pc.X+1, pc.Y)
							if tile == nil {
								tile = level.GetTileAt(pc.X, pc.Y)
								if tile.HorzTo.Level == -1 {
									playerComponent.AddMessage("You've hit the edge of the world!")
								} else {
									level.RemoveEntity(entity)
									levels[tile.HorzTo.Level].AddEntity(entity)
									pc.Level = tile.HorzTo.Level
									pc.X = tile.HorzTo.X
									pc.Y = tile.HorzTo.Y
								}
							} else if tile.Solid == true || tile.Water == true {
								playerComponent.AddMessage("You can't walk that way!")
							} else {
								pc.X++
							}
						} else {
							playerComponent.AddMessage("Something blocks you in that direction!")
						}
						dc.Direction = 0
					case "F":
						direction := playerComponent.PopCommand()
						if direction == "" {
							playerComponent.AddMessage("Wasn't given a direction to shoot!")
						} else {
							playerComponent.AddMessage("Shoot in the " + direction + " direction!")
						}
					}
				}
			}
		}
	}
}
