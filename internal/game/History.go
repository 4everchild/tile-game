package game

type History struct {
	States []Game
	Moves  []Move
}

func (h *History) GetLatest() Game {
	return h.States[len(h.States)-1]
}

func (h *History) GetPrevious() Game {
	if len(h.States) <= 1 {
		return h.States[0]
	}
	return h.GetLatest()
}

func (h *History) GetFrist() Game {
	return h.States[0]
}

func (h *History) Add(g Game, m Move) {
	h.States = append(h.States, g)
	h.Moves = append(h.Moves, m)
}

func (h *History) GetIndex(i int) Game {
	if i >= 0 && i < len(h.States) {
		return h.States[i]
	}
	return h.GetLatest()
}
