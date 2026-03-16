package game

import (
	"fmt"
	"log/slog"
)

type Move struct {
	Group uint8 `json:"group"`
	Color Color `json:"color"`
	Row   uint8 `json:"row"`
}

func (m Move) IsValid(g *Game, logger *slog.Logger) bool {

	logger.Info(fmt.Sprintf("{%d , %s , %d}", m.Group, m.Color, m.Row))
	// == means the center is selected
	if m.Group > uint8(len(g.FactoryDisplays)) {
		logger.Info("request with index out of range")
		return false
	}

	// only valid color tiles
	if !m.Color.IsTile() {
		logger.Info("color out of range for move")
		return false
	}

	//center
	if m.IsFromCenter(g) {
		switch m.Color {
		case BLUE:
			if g.Center.BLUE == 0 {
				logger.Info("tiles of this color do not exist in specified group")
				return false
			}
		case YELLOW:
			if g.Center.YELLOW == 0 {
				logger.Info("tiles of this color do not exist in specified group")
				return false
			}
		case RED:
			if g.Center.RED == 0 {
				logger.Info("tiles of this color do not exist in specified group")
				return false
			}
		case BLACK:
			if g.Center.BLACK == 0 {
				logger.Info("tiles of this color do not exist in specified group")
				return false
			}
		case GREEN:
			if g.Center.GREEN == 0 {
				logger.Info("tiles of this color do not exist in specified group")
				return false
			}
		case FIRST:
			if g.Center.FIRST == 0 {
				logger.Info("tiles of this color do not exist in specified group")
				return false
			}
		}
		// otherwise from the factory displays
	} else if g.FactoryDisplays[m.Group].CountTiles(m.Color) == 0 {
		logger.Info("tiles of this color do not exist in specified group")
		return false
	}

	//destinations check

	//for the floorline
	if m.Row == 5 {
		return true
	}

	if m.Row > 5 {
		logger.Info("row bigger than 4, there are 5 patternlines per player + the floor")
		return false
	}
	// patterline must have the same color or no color
	p := g.GetActivePlayer()
	c := g.Players[p].Patternline[m.Row].Color
	s := g.Players[p].Patternline[m.Row].Size
	if c != m.Color && c.IsTile() {
		s := fmt.Sprintf("player: %d at patternline %d already has color %s failed adding %s", p, m.Row, c.String(), m.Color.String())
		logger.Info(s)
		return false
	}

	// patternline should have some capacity
	if s >= m.Row+1 {
		return false
	}

	// corresponding tile should not be present in the wall
	// note the case for the floor is handled above
	for i := 0; i < 5; i++ {
		if g.Players[p].Wall[m.Row][i] == m.Color {
			return false
		}
	}

	return true
}

func (m *Move) IsFromCenter(g *Game) bool {
	return m.Group == uint8(len(g.FactoryDisplays))
}

//func (g* Game)
