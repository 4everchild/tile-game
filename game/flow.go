package game

//"fmt"

func (g *Game) AdvanceGame() {
	switch g.State {
	case SETUP:
		g.Setup()
	case WAITP1:
		g.State = WAITP2
	case WAITP2:
		g.State = WAITP3
	case WAITP3:
		g.State = WAITP4
	case WAITP4:
		g.State = WAITP1
	case END:
	default:
	}
}
