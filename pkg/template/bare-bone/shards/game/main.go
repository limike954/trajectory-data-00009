package main

import (
	"github.com/limike954/trajectory-data-00009/pkg/cardinal"
)

func main() {
	world, err := cardinal.NewWorld(cardinal.WorldOptions{
		TickRate:     1,
		SnapshotRate: 50,
	})
	if err != nil {
		panic(err.Error())
	}

	// Register systems
	// cardinal.RegisterSystem(world, system.ExampleSystem)

	world.StartGame()
}
