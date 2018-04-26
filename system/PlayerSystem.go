package system

import (
	"cloudlike/component"
	"cloudlike/world"
)

// PlayerSystem .
func PlayerSystem(levels []*world.Level) []*world.Level {
	for currentLevel, level := range levels {
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
								if tile.IsStairs {
									if tile.StairsTo == -1 {
										levelIndex := len(levels)
										newLevel, tX, tY := world.NewDungeon(100, 100, currentLevel, pc.X, pc.Y)
										levels = append(levels, newLevel)
										tile.StairsTo = levelIndex
										tile.StairsX = tX
										tile.StairsY = tY
									}
									level.RemoveEntity(entity)
									levels[tile.StairsTo].AddEntity(entity)
									pc.X = tile.StairsX
									pc.Y = tile.StairsY
								}
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
								if tile.IsStairs {
									if tile.StairsTo == -1 {
										levelIndex := len(levels)
										newLevel, tX, tY := world.NewDungeon(100, 100, currentLevel, pc.X, pc.Y)
										levels = append(levels, newLevel)
										tile.StairsTo = levelIndex
										tile.StairsX = tX
										tile.StairsY = tY
									}
									level.RemoveEntity(entity)
									levels[tile.StairsTo].AddEntity(entity)
									pc.X = tile.StairsX
									pc.Y = tile.StairsY
								}
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
								if tile.IsStairs {
									if tile.StairsTo == -1 {
										levelIndex := len(levels)
										newLevel, tX, tY := world.NewDungeon(100, 100, currentLevel, pc.X, pc.Y)
										levels = append(levels, newLevel)
										tile.StairsTo = levelIndex
										tile.StairsX = tX
										tile.StairsY = tY
									}
									level.RemoveEntity(entity)
									levels[tile.StairsTo].AddEntity(entity)
									pc.X = tile.StairsX
									pc.Y = tile.StairsY
								}
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
								if tile.IsStairs {
									if tile.StairsTo == -1 {
										levelIndex := len(levels)
										newLevel, tX, tY := world.NewDungeon(100, 100, currentLevel, pc.X, pc.Y)
										levels = append(levels, newLevel)
										tile.StairsTo = levelIndex
										tile.StairsX = tX
										tile.StairsY = tY
									}
									level.RemoveEntity(entity)
									levels[tile.StairsTo].AddEntity(entity)
									pc.X = tile.StairsX
									pc.Y = tile.StairsY
								}
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
	return levels
}
func triggerNewWorld(levels []*world.Level, currentLevel *world.Level, currentLevelIndex int, deltaX int, deltaY int) []*world.Level {
	levelIndex := len(levels)
	newLevel := world.NewOverworldSection(100, 100)
	levels = append(levels, newLevel)

	if deltaX > 0 {
		for y := 0; y < 100; y++ {
			tile := currentLevel.GetTileAt(99, y)
			tile.HorzTo.Level = levelIndex
			tile.HorzTo.X = 0
			tile.HorzTo.Y = y

			tile = newLevel.GetTileAt(0, y)
			tile.HorzTo.Level = currentLevelIndex
			tile.HorzTo.X = 99
			tile.HorzTo.Y = y
		}
	} else if deltaX < 0 {
		for y := 0; y < 100; y++ {
			tile := currentLevel.GetTileAt(0, y)
			tile.HorzTo.Level = levelIndex
			tile.HorzTo.X = 99
			tile.HorzTo.Y = y

			tile = newLevel.GetTileAt(99, y)
			tile.HorzTo.Level = currentLevelIndex
			tile.HorzTo.X = 0
			tile.HorzTo.Y = y
		}

	} else if deltaY < 0 {
		for x := 0; x < 100; x++ {
			tile := currentLevel.GetTileAt(x, 0)
			tile.VertTo.Level = levelIndex
			tile.VertTo.X = x
			tile.VertTo.Y = 99

			tile = newLevel.GetTileAt(x, 99)
			tile.VertTo.Level = currentLevelIndex
			tile.VertTo.X = x
			tile.VertTo.Y = 0
		}
	} else if deltaY > 0 {
		for x := 0; x < 100; x++ {
			tile := currentLevel.GetTileAt(x, 99)
			tile.VertTo.Level = levelIndex
			tile.VertTo.X = x
			tile.VertTo.Y = 0

			tile = newLevel.GetTileAt(x, 0)
			tile.VertTo.Level = currentLevelIndex
			tile.VertTo.X = x
			tile.VertTo.Y = 99
		}
	}
	return levels
}
