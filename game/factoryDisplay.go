package game

type FactoryDisplay struct {
	Tiles [4]Color `json:"tiles"`
}

func (fd *FactoryDisplay) CountTiles(c Color) int {
	count := 0

	if fd.Tiles[0] == c {
		count++
	}
	if fd.Tiles[1] == c {
		count++
	}
	if fd.Tiles[2] == c {
		count++
	}
	if fd.Tiles[3] == c {
		count++
	}

	return count

}
