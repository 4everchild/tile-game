package game

type Center struct {
	BLUE   uint8 `json:"BLUE"`
	YELLOW uint8 `json:"YELLOW"`
	RED    uint8 `json:"RED"`
	BLACK  uint8 `json:"BLACK"`
	GREEN  uint8 `json:"GREEN"`
	FIRST  uint8 `json:"FIRST"`
}

func (c *Center) HasFirst() bool {
	return c.FIRST == 1
}

func (d *Center) add(c Color, amount uint8) {
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

func (d *Center) remove(c Color, amount uint8) {
	switch c {
	case BLUE:
		d.BLUE -= amount
	case YELLOW:
		d.YELLOW -= amount
	case RED:
		d.RED -= amount
	case BLACK:
		d.BLACK -= amount
	case GREEN:
		d.GREEN -= amount
	}
}

func (d *Center) Sizeof(c Color) uint8 {
	switch c {
	case BLUE:
		return d.BLUE
	case YELLOW:
		return d.YELLOW
	case RED:
		return d.RED
	case BLACK:
		return d.BLACK
	case GREEN:
		return d.GREEN

	default:
		return 0
	}
}
