package game

import (
	"fmt"
)

func AdvanceGame(g Game) (Game, error) {
	g1 := g
	switch g1.State {
	case SETUP:
		g1.setup()
	case WAITP1:
	case WAITP2:
	case WAITP3:
	case WAITP4:
	case END:
		return g, fmt.Errorf("game ended")
	default:
		return g, fmt.Errorf("invalid game phase: %d", g.State)
	}

	return g1, nil
}
