package game

import (
	"fmt"
	"math/rand"
	"strings"
)

type Hive struct {
	Bees     []*Bee
	AliveMap map[BeeType]int
	TotalMap map[BeeType]int
}

func NewHive() *Hive {
	h := &Hive{
		Bees:     make([]*Bee, 0),
		AliveMap: make(map[BeeType]int),
		TotalMap: make(map[BeeType]int),
	}

	// Add queen (1)
	h.Bees = append(h.Bees, NewQueen())
	h.AliveMap[Queen] = 1
	h.TotalMap[Queen] = 1

	// Add workers (5)
	for i := 0; i < 5; i++ {
		h.Bees = append(h.Bees, NewWorker())
	}
	h.AliveMap[Worker] = 5
	h.TotalMap[Worker] = 5

	// Add drones (25)
	for i := 0; i < 25; i++ {
		h.Bees = append(h.Bees, NewDrone())
	}
	h.AliveMap[Drone] = 25
	h.TotalMap[Drone] = 25

	return h
}

func (h *Hive) IsAlive() bool {
	return h.AliveMap[Queen] > 0 || h.AliveMap[Worker] > 0 || h.AliveMap[Drone] > 0
}

func (h *Hive) IsQueenAlive() bool {
	return h.AliveMap[Queen] > 0
}

func (h *Hive) RandomAliveBee() *Bee {
	aliveBees := make([]*Bee, 0)
	for _, bee := range h.Bees {
		if bee.IsAlive() {
			aliveBees = append(aliveBees, bee)
		}
	}

	if len(aliveBees) == 0 {
		return nil
	}

	return aliveBees[rand.Intn(len(aliveBees))]
}

func (h *Hive) Hit() (string, bool) {
	// 20% chance to miss
	if rand.Float32() < 0.2 {
		return "Miss! You just missed the hive, better luck next time!", false
	}

	// Pick a random alive bee
	bee := h.RandomAliveBee()
	if bee == nil {
		return "The hive is already destroyed!", false
	}

	damage := bee.TakeDamage()

	// If this bee died, update counts
	if !bee.IsAlive() {
		h.AliveMap[bee.Type]--
	}

	// If queen dies, all bees die
	if bee.Type == Queen && !bee.IsAlive() {
		for _, otherBee := range h.Bees {
			if otherBee != bee && otherBee.IsAlive() {
				otherBee.HitPoints = 0
			}
		}
		h.AliveMap[Worker] = 0
		h.AliveMap[Drone] = 0
		return fmt.Sprintf("%s The Queen is dead! All other bees have died as well!", bee.AttackMessage(damage)), true
	}

	return bee.AttackMessage(damage), true
}

func (h *Hive) Sting() (string, int, bool) {
	// 15% chance to miss
	if rand.Float32() < 0.15 {
		return "Buzz! That was close! The bees just missed you!", 0, false
	}

	// Pick a random alive bee
	bee := h.RandomAliveBee()
	if bee == nil {
		return "No bees left to sting!", 0, false
	}

	return bee.StingMessage(), bee.Damage, true
}

func (h *Hive) Status() string {
	var sb strings.Builder
	sb.WriteString("Hive Status:\n")
	sb.WriteString(fmt.Sprintf("- Queen Bees: %d/%d\n", h.AliveMap[Queen], h.TotalMap[Queen]))
	sb.WriteString(fmt.Sprintf("- Worker Bees: %d/%d\n", h.AliveMap[Worker], h.TotalMap[Worker]))
	sb.WriteString(fmt.Sprintf("- Drone Bees: %d/%d\n", h.AliveMap[Drone], h.TotalMap[Drone]))

	return sb.String()
}
