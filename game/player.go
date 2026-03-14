package game

import (
	"errors"
	"fmt"
)

type Patternline struct {
	Size  uint8 `json:"size"`
	Color Color `json:"color"`
}

type Player struct {
	Patternline [5]Patternline `json:"patternline"`
	Wall        [5][5]Color    `json:"wall"`
	Floorline   [7]Color       `json:"floorline"`
	Points      uint8          `json:"points"`
}

func (p *Player) SetPatternline(index uint8, n uint8, c Color) error {
	if n > index+1 {
		errmsg := fmt.Sprintf("patternline [%d] was set at (%d) which is too much, player:\n%v\n", index, n, *p)
		return errors.New(errmsg)
	}
	p.Patternline[index].Size = n
	p.Patternline[index].Color = c
	return nil
}

// todo remove color from signature and force correct color in the wall
func (p *Player) PlaceTileWall(i, j uint8) {
	switch p.Wall[i][j] {
	case OPAQUE_BLUE:
		p.Wall[i][j] = BLUE
	case OPAQUE_BLACK:
		p.Wall[i][j] = BLACK
	case OPAQUE_RED:
		p.Wall[i][j] = RED
	case OPAQUE_YELLOW:
		p.Wall[i][j] = YELLOW
	case OPAQUE_GREEN:
		p.Wall[i][j] = GREEN
	}

}

func (p *Player) CountPoints(y, x uint8) {
	var r, l, u, d uint8 = 0, 0, 0, 0
	for i := int8(x + 1); i < 5; i++ {
		if p.Wall[y][i].IsTile() {
			r++
		} else {
			break
		}
	}
	for i := int8(x - 1); i >= 0; i-- {
		if p.Wall[y][i].IsTile() {
			l++
		} else {
			break
		}
	}
	for i := int8(y + 1); i < 5; i++ {
		if p.Wall[i][x].IsTile() {
			d++
		} else {
			break
		}
	}
	for i := int8(y - 1); i >= 0; i-- {
		if p.Wall[i][x].IsTile() {
			u++
		} else {
			break
		}
	}

	if l+r+u+d == 0 {
		p.Points += 1
		return
	}
	if l+r != 0 {
		p.Points += 1 + l + r
	}
	if u+d != 0 {
		p.Points += 1 + u + d
	}

}

func (p *Player) AddTileToFloor(c Color, g *Game) {
	var i = 0

	for i = 0; i < len(p.Floorline); i++ {
		if p.Floorline[i] == EMPTY {
			break
		}
	}

	if i >= len(p.Floorline) {
		g.Discarded.add(c, 1)
	} else {
		p.Floorline[i] = c
	}
}

func (p *Player) AddTileToPatternline(row uint8, c Color, g *Game) {
	if p.Patternline[row].Size == (row + 1) {
		p.AddTileToFloor(c, g)
	} else {
		p.Patternline[row].Size += 1
		p.Patternline[row].Color = c
	}
}

func (p *Player) PlaceFirst(g *Game) {
	if p.Floorline[6] != EMPTY {
		g.Discarded.add(p.Floorline[6], 1)
		p.Floorline[6] = FIRST
	}

	var i = 0
	for i = 0; i < len(p.Floorline); i++ {
		if p.Floorline[i] == EMPTY {
			break
		}
	}
	p.Floorline[i] = FIRST
}
