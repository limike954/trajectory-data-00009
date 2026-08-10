package multishard_test

import (
	"testing"

	"github.com/limike954/trajectory-data-00009/pkg/cardinal"
	chatsystem "github.com/limike954/trajectory-data-00009/pkg/template/multi-shard/shards/chat/system"
	gamesystem "github.com/limike954/trajectory-data-00009/pkg/template/multi-shard/shards/game/system"
)

func TestDSTGame(t *testing.T) {
	cardinal.RunDST(t, func(w *cardinal.World) {
		cardinal.RegisterSystem(w, gamesystem.PlayerSetUpdater, cardinal.WithHook(cardinal.PreUpdate))
		cardinal.RegisterSystem(w, gamesystem.PlayerSpawnSystem)
		cardinal.RegisterSystem(w, gamesystem.MovePlayerSystem)
		cardinal.RegisterSystem(w, gamesystem.PlayerLeaveSystem)
		cardinal.RegisterSystem(w, gamesystem.OnlineStatusUpdater)
	}, nil)
}

func TestDSTChat(t *testing.T) {
	cardinal.RunDST(t, func(w *cardinal.World) {
		cardinal.RegisterSystem(w, chatsystem.UserChatSystem)
	}, nil)
}
