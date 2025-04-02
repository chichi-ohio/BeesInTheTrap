package game

import (
	"strings"
	"testing"
)

func TestNewHive(t *testing.T) {
	hive := NewHive()

	// Check bee counts
	if len(hive.Bees) != 31 {
		t.Errorf("Expected 31 bees in the hive, got %d", len(hive.Bees))
	}

	if hive.AliveMap[Queen] != 1 {
		t.Errorf("Expected 1 queen bee, got %d", hive.AliveMap[Queen])
	}

	if hive.AliveMap[Worker] != 5 {
		t.Errorf("Expected 5 worker bees, got %d", hive.AliveMap[Worker])
	}

	if hive.AliveMap[Drone] != 25 {
		t.Errorf("Expected 25 drone bees, got %d", hive.AliveMap[Drone])
	}
}

func TestHiveIsAlive(t *testing.T) {
	hive := NewHive()

	if !hive.IsAlive() {
		t.Error("Expected new hive to be alive")
	}

	// Kill all bees
	for _, bee := range hive.Bees {
		bee.HitPoints = 0
	}
	hive.AliveMap[Queen] = 0
	hive.AliveMap[Worker] = 0
	hive.AliveMap[Drone] = 0

	if hive.IsAlive() {
		t.Error("Expected hive with all dead bees to be dead")
	}
}

func TestHiveIsQueenAlive(t *testing.T) {
	hive := NewHive()

	if !hive.IsQueenAlive() {
		t.Error("Expected queen to be alive in new hive")
	}

	// Find and kill queen
	for _, bee := range hive.Bees {
		if bee.Type == Queen {
			bee.HitPoints = 0
		}
	}
	hive.AliveMap[Queen] = 0

	if hive.IsQueenAlive() {
		t.Error("Expected queen to be dead after killing")
	}
}

func TestHiveRandomAliveBee(t *testing.T) {
	hive := NewHive()

	// Should return a bee
	bee := hive.RandomAliveBee()
	if bee == nil {
		t.Error("Expected to get a bee, got nil")
	}

	// Kill all bees
	for _, bee := range hive.Bees {
		bee.HitPoints = 0
	}
	hive.AliveMap[Queen] = 0
	hive.AliveMap[Worker] = 0
	hive.AliveMap[Drone] = 0

	// Should return nil
	bee = hive.RandomAliveBee()
	if bee != nil {
		t.Error("Expected nil when all bees are dead")
	}
}

func TestHiveStatus(t *testing.T) {
	hive := NewHive()

	status := hive.Status()
	if !strings.Contains(status, "Queen Bees: 1/1") {
		t.Errorf("Expected status to contain Queen count, got: %s", status)
	}

	if !strings.Contains(status, "Worker Bees: 5/5") {
		t.Errorf("Expected status to contain Worker count, got: %s", status)
	}

	if !strings.Contains(status, "Drone Bees: 25/25") {
		t.Errorf("Expected status to contain Drone count, got: %s", status)
	}
}
