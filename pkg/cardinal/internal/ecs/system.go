package ecs

import (
	"math"

	"github.com/limike954/trajectory-data-00009/pkg/assert"
	"github.com/kelindar/bitmap"
	"github.com/rotisserie/eris"
)

// SystemHook defines when a system should be executed in the update cycle.
type SystemHook uint8

const (
	// PreUpdate runs before the main update.
	PreUpdate SystemHook = 0
	// Update runs during the main update phase.
	Update SystemHook = 1
	// PostUpdate runs after the main update.
	PostUpdate SystemHook = 2
	// Init runs once during world initialization.
	Init SystemHook = 3
)

// initSystem represents a system that should be run once during world initialization.
type initSystem struct {
	name string // The name of the system
	fn   func() // Function that wraps a System
}

func RegisterSystem(world *World, name string, hook SystemHook, system func()) error {
	switch hook {
	case Init:
		world.initSystems = append(world.initSystems, initSystem{name: name, fn: system})
	case PreUpdate, Update, PostUpdate:
		world.scheduler[hook].register(name, bitmap.Bitmap{}, system)
	default:
		return eris.Errorf("invalid system hook %d", hook)
	}

	return nil
}

func RegisterSystemWithDeps(
	world *World,
	name string,
	hook SystemHook,
	system func(),
	depsComponent bitmap.Bitmap,
	depsSystemEvent bitmap.Bitmap,
) error {
	deps := depsComponent.Clone(nil)
	n := world.state.components.nextID
	assert.That(depsSystemEvent.Count()+int(n) <= math.MaxUint32-1, "system dependencies exceed max limit")
	depsSystemEvent.Range(func(x uint32) {
		deps.Set(n + x)
	})

	switch hook {
	case Init:
		world.initSystems = append(world.initSystems, initSystem{name: name, fn: system})
	case PreUpdate, Update, PostUpdate:
		world.scheduler[hook].register(name, deps, system)
	default:
		return eris.Errorf("invalid system hook %d", hook)
	}

	return nil
}
