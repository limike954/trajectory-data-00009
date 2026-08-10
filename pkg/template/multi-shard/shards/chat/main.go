package main

import (
	"github.com/limike954/trajectory-data-00009/pkg/template/multi-shard/shards/chat/system"

	"github.com/limike954/trajectory-data-00009/pkg/cardinal"
)

func main() {
	world, err := cardinal.NewWorld(cardinal.WorldOptions{
		TickRate:     20,
		SnapshotRate: 50,
	})
	if err != nil {
		panic(err.Error())
	}

	cardinal.RegisterSystem(world, system.UserChatSystem)

	world.StartGame()
}
