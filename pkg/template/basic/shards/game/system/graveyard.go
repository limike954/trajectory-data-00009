package system

import (
	"github.com/limike954/trajectory-data-00009/pkg/template/basic/shards/game/component"
	systemevent "github.com/limike954/trajectory-data-00009/pkg/template/basic/shards/game/system_event"

	"github.com/limike954/trajectory-data-00009/pkg/cardinal"
)

type GraveyardSystemState struct {
	cardinal.BaseSystemState
	PlayerDeathSystemEvents cardinal.WithSystemEventReceiver[systemevent.PlayerDeath]
	Graves                  GraveSearch
}

func GraveyardSystem(state *GraveyardSystemState) {
	for event := range state.PlayerDeathSystemEvents.Iter() {
		_, entity := state.Graves.Create()
		entity.Grave.Set(component.Gravestone{Nickname: event.Nickname})

		state.Logger().Info().Msgf("Created grave stone for player %s", event.Nickname)
	}
}
