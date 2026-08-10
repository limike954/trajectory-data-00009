package system

import (
	"github.com/limike954/trajectory-data-00009/pkg/template/basic/shards/game/component"

	"github.com/limike954/trajectory-data-00009/pkg/cardinal"
)

type PlayerSearch = cardinal.Exact[struct {
	Tag    cardinal.Ref[component.PlayerTag]
	Health cardinal.Ref[component.Health]
}]

type GraveSearch = cardinal.Exact[struct {
	Grave cardinal.Ref[component.Gravestone]
}]
