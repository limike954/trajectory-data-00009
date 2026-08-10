package lobby_test

import (
	"testing"

	"github.com/limike954/trajectory-data-00009/pkg/cardinal"
	"github.com/limike954/trajectory-data-00009/pkg/cardinal/snapshot"
	"github.com/limike954/trajectory-data-00009/pkg/lobby"
	"github.com/stretchr/testify/require"
)

func TestDST(t *testing.T) {
	cardinal.RunDST(t, func(w *cardinal.World) {
		cardinal.RegisterPlugin(w, lobby.NewPlugin(lobby.Config{}))
	}, nil)
}

func TestE2E(t *testing.T) {
	cardinal.RunE2E(t, func() *cardinal.World {
		debug := false

		world, err := cardinal.NewWorld(cardinal.WorldOptions{
			Region:              "local",
			Organization:        "organization",
			Project:             "project",
			ShardID:             "lobby",
			TickRate:            1,
			SnapshotRate:        50,
			SnapshotStorageType: snapshot.StorageTypeJetStream,
			Debug:               &debug,
		})
		require.NoError(t, err)

		cardinal.RegisterPlugin(world, lobby.NewPlugin(lobby.Config{}))

		return world
	})
}
