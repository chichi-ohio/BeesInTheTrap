package game

import (
	"testing"
)

func TestNewBees(t *testing.T) {
	// Test Queen Bee
	queen := NewQueen()
	if queen.Type != Queen {
		t.Errorf("Expected bee type to be Queen, got %v", queen.Type)
	}
	if queen.HitPoints != 100 {
		t.Errorf("Expected Queen HP to be 100, got %d", queen.HitPoints)
	}
	if queen.Damage != 10 {
		t.Errorf("Expected Queen damage to be 10, got %d", queen.Damage)
	}

	// Test Worker Bee
	worker := NewWorker()
	if worker.Type != Worker {
		t.Errorf("Expected bee type to be Worker, got %v", worker.Type)
	}
	if worker.HitPoints != 75 {
		t.Errorf("Expected Worker HP to be 75, got %d", worker.HitPoints)
	}
	if worker.Damage != 5 {
		t.Errorf("Expected Worker damage to be 5, got %d", worker.Damage)
	}

	// Test Drone Bee
	drone := NewDrone()
	if drone.Type != Drone {
		t.Errorf("Expected bee type to be Drone, got %v", drone.Type)
	}
	if drone.HitPoints != 60 {
		t.Errorf("Expected Drone HP to be 60, got %d", drone.HitPoints)
	}
	if drone.Damage != 1 {
		t.Errorf("Expected Drone damage to be 1, got %d", drone.Damage)
	}
}

func TestBeeIsAlive(t *testing.T) {
	bee := NewQueen()
	if !bee.IsAlive() {
		t.Error("Expected new bee to be alive")
	}

	bee.HitPoints = 0
	if bee.IsAlive() {
		t.Error("Expected bee with 0 HP to be dead")
	}
}

func TestBeeTakeDamage(t *testing.T) {
	// Test Queen damage
	queen := NewQueen()
	damage := queen.TakeDamage()
	if damage != 10 {
		t.Errorf("Expected queen damage to be 10, got %d", damage)
	}
	if queen.HitPoints != 90 {
		t.Errorf("Expected queen HP to be 90 after damage, got %d", queen.HitPoints)
	}

	// Test Worker damage
	worker := NewWorker()
	damage = worker.TakeDamage()
	if damage != 25 {
		t.Errorf("Expected worker damage to be 25, got %d", damage)
	}
	if worker.HitPoints != 50 {
		t.Errorf("Expected worker HP to be 50 after damage, got %d", worker.HitPoints)
	}

	// Test Drone damage
	drone := NewDrone()
	damage = drone.TakeDamage()
	if damage != 30 {
		t.Errorf("Expected drone damage to be 30, got %d", damage)
	}
	if drone.HitPoints != 30 {
		t.Errorf("Expected drone HP to be 30 after damage, got %d", drone.HitPoints)
	}

	// Test damage beyond 0
	drone.TakeDamage()
	drone.TakeDamage()
	if drone.HitPoints != 0 {
		t.Errorf("Expected drone HP to be 0 after multiple damage, got %d", drone.HitPoints)
	}
}

func TestBeeMessages(t *testing.T) {
	bee := NewQueen()

	attackMsg := bee.AttackMessage(10)
	if attackMsg == "" {
		t.Error("Expected non-empty attack message")
	}

	stingMsg := bee.StingMessage()
	if stingMsg == "" {
		t.Error("Expected non-empty sting message")
	}
}
