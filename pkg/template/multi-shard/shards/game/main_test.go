package main_test

import (
	"testing"

	"github.com/limike954/trajectory-data-00009/pkg/cardinal"
	"github.com/limike954/trajectory-data-00009/pkg/cardinal/snapshot"
	"github.com/limike954/trajectory-data-00009/pkg/template/multi-shard/shards/game/system"
	"github.com/stretchr/testify/require"
)

func TestDST(t *testing.T) {
	cardinal.RunDST(t, func(w *cardinal.World) {
		registerSystems(w)
	}, nil)
}

func TestE2E(t *testing.T) {
	cardinal.RunE2E(t, func() *cardinal.World {
		debug := false

		// Keep world setup aligned with shards/game/main.go.
		world, err := cardinal.NewWorld(cardinal.WorldOptions{
			Region:              "local",
			Organization:        "organization",
			Project:             "project",
			ShardID:             "game",
			TickRate:            20,
			SnapshotRate:        50,
			SnapshotStorageType: snapshot.StorageTypeJetStream,
			Debug:               &debug,
		})
		require.NoError(t, err)

		registerSystems(world)

		return world
	})
}

func registerSystems(w *cardinal.World) {
	cardinal.RegisterSystem(w, system.PlayerSetUpdater, cardinal.WithHook(cardinal.PreUpdate))
	cardinal.RegisterSystem(w, system.PlayerSpawnSystem)
	cardinal.RegisterSystem(w, system.MovePlayerSystem)
	cardinal.RegisterSystem(w, system.PlayerLeaveSystem)
	cardinal.RegisterSystem(w, system.OnlineStatusUpdater)
}
