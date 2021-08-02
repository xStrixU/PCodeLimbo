package player

import (
	"strings"
	"sync"
)

type PlayerManager struct {
	players sync.Map
	playerCount int64
}

func (pm *PlayerManager) AddPlayer(player *Player) {
	pm.players.Store(strings.ToLower(player.Name), player)
	pm.playerCount++
}

func (pm *PlayerManager) RemovePlayer(player *Player) {
	pm.players.Delete(strings.ToLower(player.Name))
	pm.playerCount--
}

func (pm *PlayerManager) PlayerCount() int {
	return int(pm.playerCount)
}