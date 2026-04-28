package game

//"fmt"

func (g *Game) AdvanceGame() {
	//fmt.Println(g.State)
	//fmt.Println(g.AreAllTilesPlaced())

	switch g.State {
	case SETUP:
		g.Setup()
	case WAITP1:
		if g.AreAllTilesPlaced() {
			g.State = SETUP
			g.Setup()
		} else {
			g.State = WAITP2
		}
	case WAITP2:
		if g.AreAllTilesPlaced() {
			g.State = SETUP
			g.Setup()
		} else {
			g.State = WAITP3
		}
	case WAITP3:
		if g.AreAllTilesPlaced() {
			g.State = SETUP
			g.Setup()
		} else {
			g.State = WAITP4
		}
	case WAITP4:
		if g.AreAllTilesPlaced() {
			g.State = SETUP
			g.Setup()
		} else {
			g.State = WAITP1
		}
	case END:
	default:
	}
	// TODO add endgame counting points for full: rows, columns, colors
	if g.State == END {
		for i := 0; i < len(g.Players); i++ {
			p := &g.Players[i]

			for j := 0; j < 5; j++ {
				if p.IsWallRowFull(j) {
					p.Points += 2
				}
				if p.IsWallColumnFull(j) {
					p.Points += 7
				}
				if p.IsWallColorFull(Color(j + 1)) {
					p.Points += 10
				}
			}
		}
	}

	//fmt.Println(g.State)
}
