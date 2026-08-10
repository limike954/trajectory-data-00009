package system

import (
	"github.com/limike954/trajectory-data-00009/pkg/template/basic/shards/game/component"

	"github.com/limike954/trajectory-data-00009/pkg/cardinal"
)

type RegenSystemState struct {
	cardinal.BaseSystemState
	cardinal.Contains[struct {
		cardinal.Ref[component.Health]
	}]
}

func RegenSystem(state *RegenSystemState) {
	for _, health := range state.Iter() { // Another shorthand
		health.Set(component.Health{HP: health.Get().HP + 10})
	}
}
