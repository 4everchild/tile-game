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

func (g *Game) PlaceTile(p Player, j uint8) {
	if p.Patternline[j].Size < j+1 {
		return
	}
	pci := uint8(p.Patternline[j].Color) - 1

	p.Wall[j][(pci+j)%5] = p.Patternline[j].Color
	p.CountPoints((pci+j)%5, j)
	// TODO count points
	switch pci {
	case 0:
		g.Discarded.BLUE += j
	case 1:
		g.Discarded.YELLOW += j
	case 2:
		g.Discarded.RED += j
	case 3:
		g.Discarded.BLACK += j
	case 4:
		g.Discarded.GREEN += j
	}
	p.Patternline[j].Size = 0
	p.Patternline[j].Color = EMPTY

}
