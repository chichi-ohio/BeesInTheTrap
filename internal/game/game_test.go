package game

import (
	"strings"
	"testing"
	"time"
)

func TestNewGame(t *testing.T) {
	game := NewGame()

	if game.Player == nil {
		t.Error("Expected player to be initialized")
	}

	if game.Hive == nil {
		t.Error("Expected hive to be initialized")
	}

	if game.StingCount != 0 {
		t.Errorf("Expected sting count to be 0, got %d", game.StingCount)
	}

	if game.Turn != 0 {
		t.Errorf("Expected turn to be 0, got %d", game.Turn)
	}

	if game.AutoPlayMode {
		t.Error("Expected auto play mode to be false")
	}
}

func TestGameStatus(t *testing.T) {
	game := NewGame()

	status := game.Status()
	if !strings.Contains(status, "Turn: 0") {
		t.Errorf("Expected status to contain turn info, got: %s", status)
	}

	if !strings.Contains(status, "Player Status") {
		t.Errorf("Expected status to contain player info, got: %s", status)
	}

	if !strings.Contains(status, "Hive Status") {
		t.Errorf("Expected status to contain hive info, got: %s", status)
	}
}

func TestGameIsGameOver(t *testing.T) {
	game := NewGame()

	if game.IsGameOver() {
		t.Error("Expected new game not to be over")
	}

	// Kill player
	game.Player.HitPoints = 0
	if !game.IsGameOver() {
		t.Error("Expected game to be over when player is dead")
	}

	// Reset player and kill hive
	game = NewGame()
	for _, bee := range game.Hive.Bees {
		bee.HitPoints = 0
	}
	game.Hive.AliveMap[Queen] = 0
	game.Hive.AliveMap[Worker] = 0
	game.Hive.AliveMap[Drone] = 0

	if !game.IsGameOver() {
		t.Error("Expected game to be over when hive is dead")
	}
}

func TestGamePlayerTurn(t *testing.T) {
	game := NewGame()

	initialTurn := game.Turn
	initialHitCount := game.Player.HitCount

	message := game.PlayerTurn()

	if game.Turn != initialTurn+1 {
		t.Errorf("Expected turn to increase by 1, got increase of %d", game.Turn-initialTurn)
	}

	if game.Player.HitCount != initialHitCount+1 {
		t.Errorf("Expected hit count to increase by 1, got increase of %d", game.Player.HitCount-initialHitCount)
	}

	if message == "" {
		t.Error("Expected non-empty message from player turn")
	}

	// Allow time for goroutine to complete
	time.Sleep(200 * time.Millisecond)
}

func TestGameBeeTurn(t *testing.T) {
	game := NewGame()

	message := game.BeeTurn()

	if message == "" {
		t.Error("Expected non-empty message from bee turn")
	}
}

func TestGameGetResultAndSummary(t *testing.T) {
	// Test player win
	game := NewGame()

	// Kill all bees
	for _, bee := range game.Hive.Bees {
		bee.HitPoints = 0
	}
	game.Hive.AliveMap[Queen] = 0
	game.Hive.AliveMap[Worker] = 0
	game.Hive.AliveMap[Drone] = 0

	// Set some hit count
	game.Player.HitCount = 10

	result := game.GetResult()
	if !strings.Contains(result, "Congratulations") {
		t.Errorf("Expected winning result message, got: %s", result)
	}

	summary := game.Summary()
	if !strings.Contains(summary, "WON") {
		t.Errorf("Expected winning summary, got: %s", summary)
	}

	// Test player loss
	game = NewGame()
	game.Player.HitPoints = 0
	game.StingCount = 5

	result = game.GetResult()
	if !strings.Contains(result, "Game Over") {
		t.Errorf("Expected losing result message, got: %s", result)
	}

	summary = game.Summary()
	if !strings.Contains(summary, "LOST") {
		t.Errorf("Expected losing summary, got: %s", summary)
	}
}

func TestGameToggleAutoPlay(t *testing.T) {
	game := NewGame()

	// Enable auto play
	game.ToggleAutoPlay()

	if !game.AutoPlayMode {
		t.Error("Expected auto play mode to be enabled")
	}

	// Disable auto play
	game.ToggleAutoPlay()

	if game.AutoPlayMode {
		t.Error("Expected auto play mode to be disabled")
	}
}
