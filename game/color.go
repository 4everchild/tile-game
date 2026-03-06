package game

type Color uint8

const (
	EMPTY Color = iota
	BLUE
	YELLOW
	RED
	BLACK
	GREEN
	FIRST
	OPAQUE_BLUE
	OPAQUE_YELLOW
	OPAQUE_RED
	OPAQUE_BLACK
	OPAQUE_GREEN
)

var ColorValue = map[Color]string{
	EMPTY:         "EMPTY",
	BLUE:          "BLUE",
	YELLOW:        "YELLOW",
	RED:           "RED",
	BLACK:         "BLACK",
	GREEN:         "GREEN",
	FIRST:         "FIRST",
	OPAQUE_BLUE:   "OPAQUE BLUE",
	OPAQUE_YELLOW: "OPAQUE YELLOW",
	OPAQUE_RED:    "OPAQUE RED",
	OPAQUE_BLACK:  "OPAQUE BLACK",
	OPAQUE_GREEN:  "OPAQUE GREEN",
}

func (s Color) String() string {
	return ColorValue[s]
}

var FloorLinePenalties [7]uint8 = [7]uint8{1, 1, 2, 2, 2, 3, 3}
