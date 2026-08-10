package system

import (
	"github.com/limike954/trajectory-data-00009/pkg/template/multi-shard/shards/chat/component"

	"github.com/limike954/trajectory-data-00009/pkg/cardinal"
)

type ChatSearch = cardinal.Exact[struct {
	UserTag cardinal.Ref[component.UserTag]
	Chat    cardinal.Ref[component.Chat]
}]
