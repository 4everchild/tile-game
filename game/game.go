package game

const PLAYERCOUNT = 4

type Game struct {
	Seed            XS64
	Players         [PLAYERCOUNT]Player
	FactoryDisplays [2*PLAYERCOUNT + 1]FactoryDisplay
	Center          Center    //[6]uint8
	Sack            Sack      //[5]uint8
	Discarded       Discarded //[5]uint8
	State           State
}

func NewGame(seed uint64) Game {
	var g Game // value, not pointer
	for p := range PLAYERCOUNT {
		for i := range 5 {
			if i == 0 {
				g.Players[p].Wall[i][0] = OPAQUE_BLUE
				g.Players[p].Wall[i][1] = OPAQUE_YELLOW
				g.Players[p].Wall[i][2] = OPAQUE_RED
				g.Players[p].Wall[i][3] = OPAQUE_BLACK
				g.Players[p].Wall[i][4] = OPAQUE_GREEN
			} else {
				g.Players[p].Wall[i][0] = g.Players[p].Wall[i-1][4]
				g.Players[p].Wall[i][1] = g.Players[p].Wall[i-1][0]
				g.Players[p].Wall[i][2] = g.Players[p].Wall[i-1][1]
				g.Players[p].Wall[i][3] = g.Players[p].Wall[i-1][2]
				g.Players[p].Wall[i][4] = g.Players[p].Wall[i-1][3]
			}
		}
		g.Players[p].Points = 0
	}

	g.Sack = Sack{20, 20, 20, 20, 20}
	g.State = SETUP
	if seed == 0 {
		seed = 1
	}
	g.Seed = XS64(seed)

	return g
}

func (g *Game) fillfd(i, j int) bool {
	cap := g.Sack.BLUE + g.Sack.YELLOW + g.Sack.RED + g.Sack.BLACK + g.Sack.GREEN

	if cap == 0 {
		g.Sack.BLUE += g.Discarded.BLUE
		g.Sack.YELLOW += g.Discarded.YELLOW
		g.Sack.RED += g.Discarded.RED
		g.Sack.BLACK += g.Discarded.BLACK
		g.Sack.GREEN += g.Discarded.GREEN

		g.Discarded.BLUE = 0
		g.Discarded.YELLOW = 0
		g.Discarded.RED = 0
		g.Discarded.BLACK = 0
		g.Discarded.GREEN = 0

		cap = g.Sack.BLUE + g.Sack.YELLOW + g.Sack.RED + g.Sack.BLACK + g.Sack.GREEN
	}
	if cap == 0 {
		return false
	}

	tmp := uint8(g.Seed.Step() % uint64(cap))
	c := uint8(0)
	c += g.Sack.BLUE
	if tmp < c {
		g.FactoryDisplays[i].Tiles[j] = BLUE
		g.Sack.BLUE--
		return true
	}
	c += g.Sack.YELLOW
	if tmp < c {
		g.FactoryDisplays[i].Tiles[j] = YELLOW
		g.Sack.YELLOW--
		return true
	}
	c += g.Sack.RED
	if tmp < c {
		g.FactoryDisplays[i].Tiles[j] = RED
		g.Sack.RED--
		return true
	}
	c += g.Sack.BLACK
	if tmp < c {
		g.FactoryDisplays[i].Tiles[j] = BLACK
		g.Sack.BLACK--
		return true
	}
	c += g.Sack.GREEN
	if tmp < c {
		g.FactoryDisplays[i].Tiles[j] = GREEN
		g.Sack.GREEN--
		return true
	}

	return false

}

func (g *Game) SetTile(i, j uint8, c Color) {
	switch c {
	case BLUE:
		g.Discarded.BLUE += j
	case RED:
		g.Discarded.RED += j
	case YELLOW:
		g.Discarded.YELLOW += j
	case BLACK:
		g.Discarded.BLACK += j
	case GREEN:
		g.Discarded.GREEN += j
	}

	g.Players[i].Patternline[j].Size = 0
	g.Players[i].Patternline[j].Color = EMPTY

	for k := range 5 {
		if g.Players[i].Wall[j][k].String() == "OPAQUE "+c.String() {
			g.Players[i].Wall[j][k] = c
			//score points here
		}
	}

}

func (g *Game) setup() {
	// set tile and score points score points
	for i := range uint8(PLAYERCOUNT) {
		for j := range uint8(5) {
			size := g.Players[i].Patternline[j].Size
			if size == j+1 {
				color := g.Players[i].Patternline[j].Color
				g.SetTile(i, j, color)
			}
		}
	}

	// find first player and subtract penalties from floorline
	var firstPlayer int = 0

	for i := 0; i < PLAYERCOUNT; i++ {
		for j := 0; j < 7; j++ {
			if g.Players[i].Floorline[j] == EMPTY {
				break
			}
			if g.Players[i].Floorline[j] == FIRST {
				firstPlayer = j
			}
			g.Players[i].Points -= min(FloorLinePenalties[j], g.Players[i].Points)
		}
	}

	// transition game state
	for i := 0; i < len(g.FactoryDisplays); i++ {
		g.fillfd(i, 0)
		g.fillfd(i, 1)
		g.fillfd(i, 2)
		g.fillfd(i, 3)
	}
	switch firstPlayer {
	case 0:
		g.State = WAITP1
	case 1:
		g.State = WAITP2
	case 2:
		g.State = WAITP3
	case 3:
		g.State = WAITP4
	default:
		g.State = END
	}

	// put back first tile
	g.Center.FIRST = 1
}
