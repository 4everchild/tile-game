package game

import (
	"fmt"
	"log/slog"
)

const PLAYERCOUNT = 4

type Game struct {
	Seed            XS64                              `json:"seed"`
	Players         [PLAYERCOUNT]Player               `json:"players"`
	FactoryDisplays [2*PLAYERCOUNT + 1]FactoryDisplay `json:"factorydisplays"`
	Center          Center                            `json:"center"`
	Sack            Sack                              `json:"sack"`
	Discarded       Discarded                         `json:"discarded"`
	State           State                             `json:"state"`
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

	for k := range uint8(5) {
		if g.Players[i].Wall[j][k].String() == "OPAQUE "+c.String() {
			g.Players[i].Wall[j][k] = c
			g.Players[i].CountPoints(j, k)
		}
	}

}
func (g *Game) AreAllTilesPlaced() bool {
	for i := 0; i < PLAYERCOUNT*2+1; i++ {
		// if the first tile is empty <-> all 4 are empty
		//fmt.Println(g.FactoryDisplays[i].IsEmpty())
		//fmt.Printf("%d: %t\n", i, g.FactoryDisplays[i].IsEmpty())
		if g.FactoryDisplays[i].Tiles[0] != EMPTY {
			return false
		}
	}
	if g.Center.BLUE+g.Center.RED+g.Center.YELLOW+g.Center.BLACK+g.Center.GREEN+g.Center.FIRST > 0 {
		return false
	}
	return true
}
func (g *Game) Setup() {
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
				firstPlayer = i
			}
			g.Players[i].Points -= min(FloorLinePenalties[j], g.Players[i].Points)
			g.Discarded.add(g.Players[i].Floorline[j], 1)
			g.Players[i].Floorline[j] = EMPTY
		}
	}

	// transition game state
	for i := 0; i < len(g.FactoryDisplays); i++ {
		g.fillfd(i, 0)
		g.fillfd(i, 1)
		g.fillfd(i, 2)
		g.fillfd(i, 3)
	}

	fmt.Printf("first player is: %d\n", firstPlayer)

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
	}

	// put back first tile
	g.Center.FIRST = 1

	// if row is completed// endgame
	if g.IsGameConcluded() {
		g.State = END
	}

}

func (g *Game) IsGameConcluded() bool {
	for i := 0; i < len(g.Players); i++ {
		for j := 0; j < 5; j++ {
			if g.Players[i].Wall[j][0].IsTile() &&
				g.Players[i].Wall[j][1].IsTile() &&
				g.Players[i].Wall[j][2].IsTile() &&
				g.Players[i].Wall[j][3].IsTile() &&
				g.Players[i].Wall[j][4].IsTile() {
				return true
			}
		}
	}
	return false
}

func (g *Game) GetActivePlayer() int {
	switch g.State {
	case WAITP1:
		return 0
	case WAITP2:
		return 1
	case WAITP3:
		return 2
	case WAITP4:
		return 3
	default:
		return -1
	}
}

func (g *Game) HandleMove(m Move, logger *slog.Logger) {
	//moves := g.ListAvailableMoves(&g.Players[g.GetActivePlayer()], logger)
	//fmt.Println(moves)

	g.ApplyMove(m, &g.Players[g.GetActivePlayer()])
}

func (g *Game) ApplyMove(m Move, player *Player) {
	var c Color
	//player := &g.Players[p]

	toFloor := (m.Row == 5)

	if m.IsFromCenter(g) {
		if g.Center.HasFirst() {
			player.PlaceFirst(g)
		}
		c = m.Color
		if c == FIRST {
			return
		}
		s := g.Center.Sizeof(c)
		for i := uint8(0); i < s; i++ {
			if toFloor {
				player.AddTileToFloor(c, g)
			} else {
				player.AddTileToPatternline(m.Row, c, g)
			}
		}
		g.Center.remove(c, s)
		return
	}

	for i := 0; i < 4; i++ {
		c = g.FactoryDisplays[m.Group].Tiles[i]
		if c == m.Color {
			if toFloor {
				player.AddTileToFloor(c, g)
			} else {
				player.AddTileToPatternline(m.Row, c, g)
			}
		} else {
			g.Center.add(c, 1)
		}
		g.FactoryDisplays[m.Group].Tiles[i] = EMPTY
	}

}

func (g *Game) ListAvailableMoves(p *Player, logger *slog.Logger) []Move {
	moves := make([]Move, 0)
	var move Move
	var group int
	for group = 0; group < len(g.FactoryDisplays); group++ {
		for color := BLUE; color < GREEN; color++ {
			for row := 0; row < 6; row++ {
				move = Move{uint8(group), color, uint8(row)}
				if move.IsValid(g, logger) {
					moves = append(moves, move)
				}
			}
		}
	}

	for color := BLUE; color < FIRST; color++ {
		for row := 0; row < 6; row++ {
			move = Move{uint8(group), color, uint8(row)}
			if move.IsValid(g, logger) {
				moves = append(moves, move)
			}
		}
	}

	return moves
}

func (g *Game) MakeCpuMoves(logger *slog.Logger) {
	//fmt.Println(g.AreAllTilesPlaced())
}
