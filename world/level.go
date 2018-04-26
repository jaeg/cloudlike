package world

import (
	"cloudlike/component"
	"cloudlike/entity"
	"fmt"
	"math"
	"math/rand"
)

// Level .
type Level struct {
	data                  [][]Tile
	Entities              []*entity.Entity
	width, height         int
	id                    int
	left, right, up, down int
}

type Transition struct {
	X, Y, Level int
}

//Tile .
type Tile struct {
	TileType    int
	TileIndex   int
	SpriteIndex int
	Solid       bool
	Water       bool
	floor       bool
	wall        bool
	Locked      bool
	VertTo      Transition
	HorzTo      Transition
}

func newLevel(width int, height int) (level *Level) {
	level = &Level{width: width, height: height, left: -1, right: -1, up: -1, down: -1}

	data := make([][]Tile, width, height)
	for x := 0; x < width; x++ {
		col := []Tile{}
		for y := 0; y < height; y++ {
			col = append(col, Tile{TileType: 1, TileIndex: 112, Solid: false, VertTo: Transition{Level: -1}, HorzTo: Transition{Level: -1}})
		}
		data[x] = append(data[x], col...)
	}

	level.data = data
	return
}

func NewOverworldSection(width int, height int) (level *Level) {
	fmt.Println("Creating new random level")
	level = newLevel(width, height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if rand.Intn(1000) == 0 {
				level.GetTileAt(x, y).TileType = 1
				level.GetTileAt(x, y).TileIndex = 123
			} else if rand.Intn(5) == 0 {
				level.GetTileAt(x, y).TileType = 1
				level.GetTileAt(x, y).TileIndex = 121
				level.GetTileAt(x, y).Solid = false
			} else {
				level.GetTileAt(x, y).TileType = 2
				level.GetTileAt(x, y).TileIndex = 122
			}
		}
	}

	//Generate Flower Medows
	for i := 0; i < 20; i++ {
		x := getRandom(1, width)
		y := getRandom(1, height)

		level.createCluster(x, y, 10, 123, 0, false, false)
	}

	//Generate Water
	for i := 0; i < 10; i++ {
		x := getRandom(1, width)
		y := getRandom(1, height)

		level.createCluster(x, y, 100, 181, 0, false, true)
	}

	return
}

func (level *Level) GetTileAt(x int, y int) (tile *Tile) {
	if x < level.width && y < level.height && x >= 0 && y >= 0 {
		tile = &level.data[x][y]
	}
	return
}

//Get's the view frustum with the player in the center
func (level *Level) GetView(aX int, aY int, width int, height int, blind bool) (data [][]*Tile) {
	left := aX - width/2
	right := aX + width/2
	up := aY - height/2
	down := aY + height/2

	data = make([][]*Tile, width+1-width%2)

	cX := 0
	for x := left; x <= right; x++ {
		col := []*Tile{}
		for y := up; y <= down; y++ {
			currentTile := level.GetTileAt(x, y)
			if blind {
				if y < aY-height/4 || y > aY+height/4 || x > aX+width/4 || x < aX-width/4 {
					currentTile = nil
				}
			}

			if currentTile != nil {
				if los(aX, aY, x, y, level) == false {
					currentTile = nil
				}
			}

			col = append(col, currentTile)
		}
		data[cX] = append(data[cX], col...)
		cX++
	}
	return
}

func (level *Level) GetEntityAt(x int, y int) (entity *entity.Entity) {
	for i := 0; i < len(level.Entities); i++ {
		entity = level.Entities[i]
		if entity.HasComponent("PositionComponent") {
			pc := entity.GetComponent("PositionComponent").(*component.PositionComponent)
			if pc.X == x && pc.Y == y {
				return
			}
		}
	}
	entity = nil
	return
}

func (level *Level) GetSolidEntityAt(x int, y int) (entity *entity.Entity) {
	for i := 0; i < len(level.Entities); i++ {
		entity = level.Entities[i]
		if entity.HasComponent("PositionComponent") {
			if entity.HasComponent("SolidComponent") {
				pc := entity.GetComponent("PositionComponent").(*component.PositionComponent)
				if pc.X == x && pc.Y == y {
					return
				}
			}
		}
	}
	entity = nil
	return
}

func (level *Level) AddEntity(entity *entity.Entity) {
	level.Entities = append(level.Entities, entity)
}

func (level *Level) RemoveEntity(entity *entity.Entity) {
	for i := 0; i < len(level.Entities); i++ {
		if level.Entities[i] == entity {
			level.Entities = append(level.Entities[:i], level.Entities[i+1:]...)

		}
	}
}

func (level *Level) createCluster(x int, y int, size int, tileIndex int, spriteIndex int, solid bool, water bool) {
	for i := 0; i < 200; i++ {
		n := getRandom(1, 6)
		e := getRandom(1, 6)
		s := getRandom(1, 6)
		w := getRandom(1, 6)

		if n == 1 {
			x += 1
		}

		if s == 1 {
			x--
		}

		if e == 1 {
			y++
		}

		if w == 1 {
			y--
		}

		if level.GetTileAt(x, y) != nil {
			tile := level.GetTileAt(x, y)
			tile.SpriteIndex = spriteIndex
			tile.TileIndex = tileIndex
			tile.Water = water
			tile.Solid = solid
		}
	}
}

func getRandom(low int, high int) int {
	return (rand.Intn((high - low))) + low
}

func Sgn(a int) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return +1
	}
	return 0
}

//Ported from http://www.roguebasin.com/index.php?title=Simple_Line_of_Sight
func los(pX int, pY int, tX int, tY int, level *Level) bool {
	deltaX := pX - tX
	deltaY := pY - tY

	absDeltaX := math.Abs(float64(deltaX))
	absDeltaY := math.Abs(float64(deltaY))

	signX := Sgn(deltaX)
	signY := Sgn(deltaY)

	if absDeltaX > absDeltaY {
		t := absDeltaY*2 - absDeltaX
		for {
			if t >= 0 {
				tY += signY
				t -= absDeltaX * 2
			}

			tX += signX
			t += absDeltaY * 2

			if tX == pX && tY == pY {
				return true
			}
			if level.GetTileAt(tX, tY).Solid == true {
				break
			}
		}
		return false
	}

	t := absDeltaX*2 - absDeltaY

	for {
		if t >= 0 {
			tX += signX
			t -= absDeltaY * 2
		}
		tY += signY
		t += absDeltaX * 2
		if tX == pX && tY == pY {
			return true
		}

		if level.GetTileAt(tX, tY).Solid == true {
			break
		}
	}

	return false

}
