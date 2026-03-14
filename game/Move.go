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
		// otherwise the factory displays
	} else if g.FactoryDisplays[m.Group].CountTiles(m.Color) == 0 {
		logger.Info("tiles of this color do not exist in specified group")
		return false
	}

	if m.Row > 4 {
		logger.Info("row bigger than 4, there are 5 patternlines per player")
		return false
	}

	p := g.GetActivePlayer()
	c := g.Players[p].Patternline[m.Row].Color
	if c != m.Color && c.IsTile() {
		s := fmt.Sprintf("player: %d at patternline %d already has color %s failed adding %s", p, m.Row, c.String(), m.Color.String())
		logger.Info(s)
		return false
	}

	return true
}

func (m *Move) IsFromCenter(g *Game) bool {
	return m.Group == uint8(len(g.FactoryDisplays))
}

//func (g* Game)
