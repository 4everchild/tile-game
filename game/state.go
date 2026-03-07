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
	SETUP:  "SETUP",
	END:    "END",
	WAITP1: "P1",
	WAITP2: "P2",
	WAITP3: "P3",
	WAITP4: "P4",
}

func (s State) String() string {
	return StateName[s]
}
