package game

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
	defer g.Center.remove(FIRST, 1)
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
	if i == 7 {
		g.Discarded.add(p.Floorline[6], 1)
		p.Floorline[6] = FIRST
	} else {
		p.Floorline[i] = FIRST
	}
}

func (p *Player) IsWallRowFull(r int) bool {
	if p.Wall[r][0].IsTile() &&
		p.Wall[r][1].IsTile() &&
		p.Wall[r][2].IsTile() &&
		p.Wall[r][3].IsTile() &&
		p.Wall[r][4].IsTile() {
		//fmt.Printf("row %d for player %v", r, p.Points)
		return true
	}
	return false
}

func (p *Player) IsWallColumnFull(c int) bool {
	if p.Wall[0][c].IsTile() &&
		p.Wall[1][c].IsTile() &&
		p.Wall[2][c].IsTile() &&
		p.Wall[3][c].IsTile() &&
		p.Wall[4][c].IsTile() {
		//fmt.Printf("column %d for player %v", c, p.Points)
		return true
	}
	return false
}

func (p *Player) IsWallColorFull(c Color) bool {
	switch c {
	case BLUE:
		return p.Wall[0][0].IsTile() && p.Wall[1][1].IsTile() && p.Wall[2][2].IsTile() && p.Wall[3][3].IsTile() && p.Wall[4][4].IsTile()
	case YELLOW:
		return p.Wall[0][1].IsTile() && p.Wall[1][2].IsTile() && p.Wall[2][3].IsTile() && p.Wall[3][4].IsTile() && p.Wall[4][0].IsTile()
	case RED:
		return p.Wall[0][2].IsTile() && p.Wall[1][3].IsTile() && p.Wall[2][4].IsTile() && p.Wall[3][0].IsTile() && p.Wall[4][1].IsTile()
	case BLACK:
		return p.Wall[0][3].IsTile() && p.Wall[1][4].IsTile() && p.Wall[2][0].IsTile() && p.Wall[3][1].IsTile() && p.Wall[4][2].IsTile()
	case GREEN:
		return p.Wall[0][4].IsTile() && p.Wall[1][0].IsTile() && p.Wall[2][1].IsTile() && p.Wall[3][2].IsTile() && p.Wall[4][3].IsTile()
	default:
		return false
	}
}
