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
	IsStairs    bool
	StairsTo    int
	StairsX     int
	StairsY     int
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
			col = append(col, Tile{TileType: 1, TileIndex: 112, Solid: false, StairsTo: -1, IsStairs: false, VertTo: Transition{Level: -1}, HorzTo: Transition{Level: -1}})
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
				level.GetTileAt(x, y).TileIndex = 13
				level.GetTileAt(x, y).StairsTo = -1
				level.GetTileAt(x, y).IsStairs = true
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

	//create a town
	level.buildRecursiveRoom(0, 0, 20, 20, 9)

	return
}

func NewDungeon(width int, height int, stairsUpTo int, fromX int, fromY int) (level *Level, pX int, pY int) {
	fmt.Println("Creating dungeon")
	level = newLevel(width, height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if rand.Intn(1000) == 0 {
				level.GetTileAt(x, y).TileType = 1
				level.GetTileAt(x, y).TileIndex = 12
				level.GetTileAt(x, y).StairsTo = -1
				level.GetTileAt(x, y).IsStairs = true
			} else if rand.Intn(5) == 0 {
				level.GetTileAt(x, y).TileType = 1
				level.GetTileAt(x, y).TileIndex = 8
				level.GetTileAt(x, y).Solid = false
			} else {
				level.GetTileAt(x, y).TileType = 2
				level.GetTileAt(x, y).TileIndex = 7
			}
		}
	}

	pX = getRandom(0, width)
	pY = getRandom(0, height)
	level.GetTileAt(pX, pY).IsStairs = true
	level.GetTileAt(pX, pY).StairsTo = stairsUpTo
	level.GetTileAt(pX, pY).StairsX = fromX
	level.GetTileAt(pX, pY).StairsY = fromY
	level.GetTileAt(pX, pY).TileIndex = 13
	level.GetTileAt(pX, pY).Solid = false

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
	if low == high {
		return low
	}
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

// Ported from here: https://github.com/tome2/tome2/blob/master/src/generate.cc
func (level *Level) buildRecursiveRoom(x1 int, y1 int, x2 int, y2 int, power int) {
	xSize := x2 - x1
	ySize := y2 - y1

	if xSize < 0 || ySize < 0 {
		return
	}

	choice := 0
	if power > 3 && xSize > 12 && ySize > 12 {
		choice = 1
	} else {
		if power < 10 {
			if getRandom(0, 10) > 2 && xSize < 8 && ySize < 8 {
				choice = 4
			} else {
				choice = getRandom(2, 4)
				fmt.Println(choice)
			}
		}
	}

	if choice == 1 {
		//Outer walls
		for x := x1; x <= x2; x++ {
			level.GetTileAt(x, y1).TileIndex = 0
			level.GetTileAt(x, y2).TileIndex = 0
		}

		for y := y1 + 1; y < y2; y++ {
			level.GetTileAt(x1, y).TileIndex = 6
			level.GetTileAt(x2, y).TileIndex = 6
		}

		if getRandom(0, 2) == 0 {
			y := getRandom(0, ySize) + y1
			level.GetTileAt(x1, y).TileIndex = 121
			level.GetTileAt(x2, y).TileIndex = 121
		} else {

			x := getRandom(0, xSize) + x1
			level.GetTileAt(x, y1).TileIndex = 121
			level.GetTileAt(x, y2).TileIndex = 121
		}

		//Size of keep
		t1 := getRandom(0, ySize/3) + y1
		t2 := y2 - getRandom(0, ySize/3)
		t3 := getRandom(0, xSize/3) + x1
		t4 := x2 - getRandom(0, xSize/3)

		//Above and below
		level.buildRecursiveRoom(x1+1, y1+1, x2-1, t1, power+1)
		level.buildRecursiveRoom(x1+1, t2, x2-1, y2, power+1)

		//Left and right
		level.buildRecursiveRoom(x1+1, t1+1, t3, t2-1, power+3)
		level.buildRecursiveRoom(t4, t1+1, x2-1, t2-1, power+3)

		x1 = t3
		x2 = t4
		y1 = t1
		y2 = t2
		xSize = x2 - x1
		ySize = y2 - y1
		power += 2
	}

	if choice == 4 || choice == 1 {
		if xSize < 3 || ySize < 3 {
			for y := y1; y < y2; y++ {
				for x := x1; x < x2; x++ {
					//level.GetTileAt(x, y).TileIndex = 0
				}
			}

			return
		}

		//make outside walls
		for x := x1 + 1; x <= x2-1; x++ {
			level.GetTileAt(x, y1+1).TileIndex = 0
			level.GetTileAt(x, y2-1).TileIndex = 0
		}

		for y := y1 + 1; y < y2-1; y++ {
			level.GetTileAt(x1+1, y).TileIndex = 6
			level.GetTileAt(x2-1, y).TileIndex = 6
		}

		//Make door
		y := getRandom(0, ySize-3) + y1 + 1
		if getRandom(0, 2) == 0 {
			level.GetTileAt(x1+1, y).TileIndex = 121
		} else {
			level.GetTileAt(x2-1, y).TileIndex = 121
		}

		level.buildRecursiveRoom(x1+2, y1+2, x2-2, y2-2, power+3)
	}

	if choice == 2 {
		if xSize < 3 {
			for y := y1; y < y2; y++ {
				for x := x1; x < x2; x++ {
					//level.GetTileAt(x, y).TileIndex = 6
				}
			}

			return
		}

		t1 := getRandom(0, xSize-2) + x1 + 1
		level.buildRecursiveRoom(x1, y1, t1, y2, power-2)
		level.buildRecursiveRoom(t1+1, y1, x2, y2, power-2)
	}

	if choice == 3 {
		if ySize < 3 {
			for y := y1; y < y2; y++ {
				for x := x1; x < x2; x++ {
					//level.GetTileAt(x, y).TileIndex = 6
				}
			}

			return
		}

		t1 := getRandom(0, ySize-2) + y1 + 1
		level.buildRecursiveRoom(x1, y1, x2, t1, power-2)
		level.buildRecursiveRoom(x1, t1+1, x2, y2, power-2)
	}
}
