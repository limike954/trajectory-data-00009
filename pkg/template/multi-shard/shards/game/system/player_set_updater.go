package system

import (
	"github.com/limike954/trajectory-data-00009/pkg/cardinal"
)

type PlayerSetUpdaterState struct {
	cardinal.BaseSystemState
	Players PlayerSearch
}

// PlayerSetUpdater updates the playerSet with all players in the world state.
func PlayerSetUpdater(state *PlayerSetUpdaterState) {
	playerSet.Clear()
	for _, player := range state.Players.Iter() {
		playerSet.Add(player.Tag.Get().ArgusAuthID)
	}
}
