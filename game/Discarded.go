package game

type Discarded struct {
	BLUE   uint8 `json:"BLUE"`
	YELLOW uint8 `json:"YELLOW"`
	RED    uint8 `json:"RED"`
	BLACK  uint8 `json:"BLACK"`
	GREEN  uint8 `json:"GREEN"`
}

func (d *Discarded) add(c Color, amount uint8) {
	switch c {
	case BLUE:
		d.BLUE += amount
	case YELLOW:
		d.YELLOW += amount
	case RED:
		d.RED += amount
	case BLACK:
		d.BLACK += amount
	case GREEN:
		d.GREEN += amount
	}

}
