package game

import "fmt"

//"fmt"

func (g *Game) AdvanceGame() {
	fmt.Println(g.State)
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
	fmt.Println(g.State)
}
