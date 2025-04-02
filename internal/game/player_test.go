package game

import (
	"strings"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	player := NewPlayer()

	if player.HitPoints != 100 {
		t.Errorf("Expected player to have 100 HP, got %d", player.HitPoints)
	}

	if player.MaxHP != 100 {
		t.Errorf("Expected player to have 100 max HP, got %d", player.MaxHP)
	}

	if player.HitCount != 0 {
		t.Errorf("Expected player to have 0 hit count, got %d", player.HitCount)
	}
}

func TestPlayerIsAlive(t *testing.T) {
	player := NewPlayer()

	if !player.IsAlive() {
		t.Error("Expected new player to be alive")
	}

	player.HitPoints = 0

	if player.IsAlive() {
		t.Error("Expected player with 0 HP to be dead")
	}
}

func TestPlayerTakeDamage(t *testing.T) {
	player := NewPlayer()

	// Take 10 damage
	player.TakeDamage(10)
	if player.HitPoints != 90 {
		t.Errorf("Expected 90 HP after taking 10 damage, got %d", player.HitPoints)
	}

	// Take damage that would reduce HP below 0
	player.TakeDamage(100)
	if player.HitPoints != 0 {
		t.Errorf("Expected 0 HP after taking excessive damage, got %d", player.HitPoints)
	}
}

func TestPlayerHit(t *testing.T) {
	player := NewPlayer()

	// Make 3 hits
	for i := 0; i < 3; i++ {
		player.Hit()
	}

	if player.HitCount != 3 {
		t.Errorf("Expected hit count to be 3, got %d", player.HitCount)
	}
}

func TestPlayerStatus(t *testing.T) {
	player := NewPlayer()
	player.HitCount = 5

	status := player.Status()

	if !strings.Contains(status, "100/100 HP") {
		t.Errorf("Expected status to contain HP info, got: %s", status)
	}

	if !strings.Contains(status, "Hits: 5") {
		t.Errorf("Expected status to contain hit count, got: %s", status)
	}
}
