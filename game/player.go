package game

import (
	"errors"
	"fmt"
)

type Patternline struct {
	Size  uint8
	Color Color
}

type Player struct {
	Patternline [5]Patternline
	Wall        [5][5]Color
	Floorline   [7]Color
	Points      uint8
}

func (p *Player) SetPatternline(index uint8, n uint8, c Color) error {
	if n >= index {
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
	fmt.Printf("counting points for placements in (%d,%d)\n", x, y)
	var r, l, u, d uint8 = 0, 0, 0, 0
	//p.Points = r + l + u + d
	for i := int8(x + 1); i < 5; i++ {
		if p.Wall[y][i].IsTile() {
			fmt.Println("r")
			r++
		} else {
			break
		}
	}
	for i := int8(x - 1); i >= 0; i-- {
		if p.Wall[y][i].IsTile() {
			fmt.Println("l")
			l++
		} else {
			break
		}
	}
	for i := int8(y + 1); i < 5; i++ {
		if p.Wall[i][x].IsTile() {
			fmt.Println("d")
			d++
		} else {
			break
		}
	}
	for i := int8(y - 1); i >= 0; i-- {
		if p.Wall[i][x].IsTile() {
			fmt.Println("u")
			u++
		} else {
			break
		}
	}
	fmt.Printf("r: %d\nl: %d\nu: %d\nd: %d\n", r, l, u, d)

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
