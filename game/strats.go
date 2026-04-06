package game

import (
	"fmt"
	"log/slog"
)

func (g *Game) MakeRandomMove(logger *slog.Logger) Move {
	p := &g.Players[g.GetActivePlayer()]

	moves := g.ListAvailableMoves(p, logger)
	//fmt.Println(moves)
	n := g.Seed.Step() % uint64(len(moves))
	//fmt.Println(n)
	//fmt.Println(moves[n])

	g.ApplyMove(moves[n], p)
	return moves[n]
}

func (g *Game) MakeRandomSensibleMove(logger *slog.Logger) Move {
	p := &g.Players[g.GetActivePlayer()]

	moves := g.ListAvailablePlMoves(p, logger)
	if len(moves) == 0 {
		moves = g.ListAvailableMoves(p, logger)
	}

	//fmt.Println(moves)
	n := g.Seed.Step() % uint64(len(moves))

	//fmt.Println(n)
	//fmt.Println(moves[n])

	g.ApplyMove(moves[n], p)
	return moves[n]
}

func (g *Game) MakeBestMove(logger *slog.Logger, score func(Move) float32) {
	p := &g.Players[g.GetActivePlayer()]

	moves := g.ListAvailableMoves(p, logger)

	scores := make([]float32, len(moves))
	for i := 0; i < len(moves); i++ {
		scores[i] = score(moves[i])
		fmt.Printf("%+v score: %f", moves[i], scores[i])
	}

}
