package game

import (
	"errors"
	"sync"
)

type GameManager struct {
	mu      sync.RWMutex       // protects the map
	History map[uint64]History // map[id]*History
	nextID  uint64             // auto-incrementing ID
}

func NewGameManager() *GameManager {
	return &GameManager{
		History: make(map[uint64]History),
		nextID:  1,
	}
}

func (m *GameManager) GenerateId() uint64 {
	m.mu.Lock()
	defer m.mu.Unlock()
	id := m.nextID
	m.nextID++
	return id
}

func (m *GameManager) CreateGame() uint64 {
	id := m.GenerateId()
	game := NewGame(id)
	game.AdvanceGame()

	h := m.History[id]
	h.States = append(h.States, game)
	m.History[id] = h

	return id

}

func (m *GameManager) GetLatestGame(id uint64) (Game, error) {
	h, ok := m.History[id]

	if ok {
		return h.GetLatest(), nil
	} else {
		return Game{}, errors.New("id not found")
	}
}

func (m *GameManager) SetLatestGame(id uint64, g Game) error {
	h, ok := m.History[id]

	if ok {
		h.States[len(h.States)-1] = g
		return nil
	} else {
		return errors.New("id not found")
	}
}
