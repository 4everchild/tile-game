package game

import (
	"errors"
	"sync"
)

type GameManager struct {
	mu     sync.RWMutex    // protects the map
	Games  map[uint64]Game // map[id]*Game
	nextID uint64          // auto-incrementing ID
}

func NewGameManager() *GameManager {
	return &GameManager{
		Games:  make(map[uint64]Game),
		nextID: 1,
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
	game.Setup()
	m.Games[id] = game

	return id

}

func (m *GameManager) GetGame(id uint64) (Game, error) {
	g, ok := m.Games[id]

	if ok {
		return g, nil
	} else {
		return Game{}, errors.New("id not found")
	}
}
