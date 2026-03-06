package game

type State uint8

const (
	SETUP State = iota
	END
	WAITP1
	WAITP2
	WAITP3
	WAITP4
)

var StateName = map[State]string{
	SETUP:  "### setup ###",
	END:    "### GAME ENDED ###",
	WAITP1: "### waiting p1 ###",
	WAITP2: "### waiting p2 ###",
	WAITP3: "### waiting p3 ###",
	WAITP4: "### waiting p4 ###",
}

func (s State) String() string {
	return StateName[s]
}
