package ecs

import (
	"time"

	"slices"

	"github.com/kelindar/bitmap"
)

// systemMetadata contains the metadata for a system.
type systemMetadata struct {
	name string        // The name of the system
	deps bitmap.Bitmap // Bitmap of system dependencies (components + system events)
	fn   func()        // Function that wraps a System
}

// systemScheduler manages deterministic sequential execution of systems.
type systemScheduler struct {
	systemHook SystemHook       // The execution hook ("pre", "update", "post")
	systems    []systemMetadata // The systems to run
	tier0      []int            // The first execution tier, retained for schedule introspection tests
	graph      map[int][]int    // Mapping of systems -> systems that depend on it
	// onSystemRun is an optional callback for debug performance metrics.
	// When set, it is called after each system execution with the system name, system hook,
	// and start/end times. Duration is derived from monotonic clock (time.Since);
	onSystemRun func(name string, hook SystemHook, startTime, endTime time.Time)
}

// newSystemScheduler creates a new system scheduler.
func newSystemScheduler() systemScheduler {
	return systemScheduler{
		systems: make([]systemMetadata, 0),
		tier0:   make([]int, 0),
		graph:   make(map[int][]int),
	}
}

// register registers a system with the scheduler.
func (s *systemScheduler) register(name string, systemDep bitmap.Bitmap, systemFn func()) {
	s.systems = append(s.systems, systemMetadata{name: name, deps: systemDep, fn: systemFn})
}

// Run executes the systems in registration order.
func (s *systemScheduler) Run() {
	// Single monotonic base for this run
	schedulerStartTime := time.Now()

	for _, system := range s.systems {
		if s.onSystemRun != nil { // Set by the debug module to collect performance metrics.
			systemStartTime := schedulerStartTime.Add(time.Since(schedulerStartTime))
			system.fn()
			systemEndTime := systemStartTime.Add(time.Since(systemStartTime))
			s.onSystemRun(system.name, s.systemHook, systemStartTime, systemEndTime)
		} else {
			system.fn()
		}
	}
}

// SystemInfo describes a system and its dependents for external introspection.
type SystemInfo struct {
	ID         int
	Name       string
	Dependents []int // IDs of systems that depend on this one (forward edges).
}

// ScheduleInfo describes the dependency graph for one execution phase.
type ScheduleInfo struct {
	Hook    SystemHook
	Systems []SystemInfo
}

// scheduleInfo returns the dependency graph for introspection.
func (s *systemScheduler) scheduleInfo() ScheduleInfo {
	systems := make([]SystemInfo, len(s.systems))
	for i, sys := range s.systems {
		systems[i] = SystemInfo{
			ID:         i,
			Name:       sys.name,
			Dependents: s.graph[i],
		}
	}
	return ScheduleInfo{Hook: s.systemHook, Systems: systems}
}

// createSchedule initializes schedule metadata for introspection.
// Must be called after all systems are registered and before the first Run.
func (s *systemScheduler) createSchedule() {
	graph, indegree := buildDependencyGraph(s.systems)
	s.graph = graph
	s.tier0 = getFirstTier(s.systems, indegree)
}

// buildDependencyGraph creates a directed acyclic graph (DAG) of system dependencies
// based on their shared component access patterns. The graph is retained for debug
// introspection; execution is always sequential in registration order.
func buildDependencyGraph(systems []systemMetadata) (map[int][]int, map[int]int) {
	graph := make(map[int][]int, len(systems))
	indegree := make(map[int]int, len(systems))

	for i := range systems {
		graph[i] = nil
	}

	for systemA := range len(systems) - 1 {
		for systemB := systemA + 1; systemB < len(systems); systemB++ {
			depsA := systems[systemA].deps
			depsB := systems[systemB].deps

			var deps []uint32
			depsA.Range(func(x uint32) {
				deps = append(deps, x)
			})

			// Check if systemB depends on systemA.
			if slices.ContainsFunc(deps, depsB.Contains) {
				graph[systemA] = append(graph[systemA], systemB)
				indegree[systemB]++
			}
		}
	}

	return graph, indegree
}

// getFirstTier returns the list of systems without any dependencies.
func getFirstTier(systems []systemMetadata, indegree map[int]int) []int {
	var currentTier []int
	for systemID := range systems {
		if indegree[systemID] == 0 {
			currentTier = append(currentTier, systemID)
		}
	}
	return currentTier
}
