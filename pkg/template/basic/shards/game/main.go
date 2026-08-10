package main

import (
	"github.com/limike954/trajectory-data-00009/pkg/cardinal/snapshot"
	"github.com/limike954/trajectory-data-00009/pkg/template/basic/shards/game/system"

	"github.com/limike954/trajectory-data-00009/pkg/cardinal"
)

func main() {
	world, err := cardinal.NewWorld(cardinal.WorldOptions{
		TickRate:            1,
		SnapshotRate:        50,
		SnapshotStorageType: snapshot.StorageTypeJetStream,
	})
	if err != nil {
		panic(err.Error())
	}

	cardinal.RegisterSystem(world, system.PlayerSpawnerSystem, cardinal.WithHook(cardinal.Init))

	cardinal.RegisterSystem(world, system.CreatePlayerSystem)
	cardinal.RegisterSystem(world, system.RegenSystem)
	cardinal.RegisterSystem(world, system.AttackPlayerSystem)
	cardinal.RegisterSystem(world, system.GraveyardSystem)
	cardinal.RegisterSystem(world, system.CallExternalSystem)

	world.StartGame()
}
