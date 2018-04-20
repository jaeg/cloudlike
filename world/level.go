package world

import (
	"fmt"
	"math/rand"
)

// Level .
type Level struct {
	data                  [][]Tile
	width, height         int
	id                    int
	left, right, up, down int
}

//Tile .
type Tile struct {
	TileType    int
	SpriteIndex int
	solid       bool
	floor       bool
	wall        bool
	Locked      bool
}

func newLevel(width int, height int) (level *Level) {
	level = &Level{width: width, height: height, left: -1, right: -1, up: -1, down: -1}

	data := make([][]Tile, width, height)
	for x := 0; x < width; x++ {
		col := []Tile{}
		for y := 0; y < height; y++ {
			col = append(col, Tile{TileType: 1, SpriteIndex: 112, solid: false})
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
			if rand.Intn(2) == 0 {
				level.getTileAt(x, y).TileType = 1
				level.getTileAt(x, y).SpriteIndex = 112
			} else {
				level.getTileAt(x, y).TileType = 2
				level.getTileAt(x, y).SpriteIndex = 96
			}
		}
	}

	return
}

func (level *Level) getTileAt(x int, y int) (tile *Tile) {
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
			currentTile := level.getTileAt(x, y)
			if blind {
				if y < aY-height/4 || y > aY+height/4 || x > aX+width/4 || x < aX-width/4 {
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
