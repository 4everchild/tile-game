package game

import (
	"log/slog"
)

func (g *Game) MakeRandomMove(logger *slog.Logger) {
	p := &g.Players[g.GetActivePlayer()]

	moves := g.ListAvailableMoves(p, logger)
	//fmt.Println(moves)
	n := g.Seed.Step() % uint64(len(moves))
	//fmt.Println(n)
	//fmt.Println(moves[n])

	g.ApplyMove(moves[n], p)
}
