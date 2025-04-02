package game

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Game struct {
	Player       *Player
	Hive         *Hive
	StingCount   int
	Turn         int
	AutoPlayMode bool
	AutoPlayWg   *sync.WaitGroup
	StopChan     chan struct{}
}

func NewGame() *Game {
	// Initialize random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	return &Game{
		Player:       NewPlayer(),
		Hive:         NewHive(),
		StingCount:   0,
		Turn:         0,
		AutoPlayMode: false,
		AutoPlayWg:   &sync.WaitGroup{},
		StopChan:     make(chan struct{}),
	}
}

func (g *Game) Status() string {
	return fmt.Sprintf(`Turn: %d

Player Status:
  Health: %d/%d HP
  Hits: %d

Hive Status:
  Queen Bees:  %d/%d
  Worker Bees: %d/%d
  Drone Bees:  %d/%d`,
		g.Turn,
		g.Player.HitPoints, g.Player.MaxHP, g.Player.HitCount,
		g.Hive.AliveMap[Queen], g.Hive.TotalMap[Queen],
		g.Hive.AliveMap[Worker], g.Hive.TotalMap[Worker],
		g.Hive.AliveMap[Drone], g.Hive.TotalMap[Drone])
}

func (g *Game) IsGameOver() bool {
	return !g.Player.IsAlive() || !g.Hive.IsAlive()
}

func (g *Game) PlayerTurn() string {
	g.Turn++
	g.Player.Hit()

	message, hit := g.Hive.Hit()
	if hit {
		// Process successful hit in a goroutine to demonstrate concurrency
		go func() {
			time.Sleep(100 * time.Millisecond) // Simulate some processing
			// Check if queen died
			if !g.Hive.IsQueenAlive() {
				g.Hive.AliveMap[Worker] = 0
				g.Hive.AliveMap[Drone] = 0
			}
		}()
	}

	return message
}

func (g *Game) BeeTurn() string {
	message, damage, hit := g.Hive.Sting()
	if hit {
		g.StingCount++
		g.Player.TakeDamage(damage)
	}

	return message
}

func (g *Game) GetResult() string {
	if g.Player.IsAlive() {
		return fmt.Sprintf("Congratulations! You destroyed the hive in %d hits!", g.Player.HitCount)
	} else {
		return fmt.Sprintf("Game Over! The hive killed you with %d stings.", g.StingCount)
	}
}

func (g *Game) Summary() string {
	var result string

	if g.Player.IsAlive() {
		result = fmt.Sprintf("You WON! You destroyed the hive with %d HP remaining.\n", g.Player.HitPoints)
		result += fmt.Sprintf("It took you %d hits to destroy the hive.\n", g.Player.HitCount)
	} else {
		result = fmt.Sprintf("You LOST! The hive killed you after %d turns.\n", g.Turn)
		result += fmt.Sprintf("The bees stung you %d times before you were defeated.\n", g.StingCount)
		result += "Remaining bees:\n"
		result += g.Hive.Status()
	}

	return result
}

func (g *Game) ToggleAutoPlay() {
	if g.AutoPlayMode {
		// If already in auto play, stop it
		close(g.StopChan)
		g.AutoPlayWg.Wait()

		// Reset channels for next time
		g.StopChan = make(chan struct{})
		g.AutoPlayMode = false
		return
	}

	// Start auto play
	g.AutoPlayMode = true
	g.AutoPlayWg.Add(1)

	// Run auto play in a goroutine
	go func() {
		defer g.AutoPlayWg.Done()

		for !g.IsGameOver() {
			select {
			case <-g.StopChan:
				return
			default:
				g.PlayerTurn()

				// Exit if game over after player turn
				if g.IsGameOver() {
					return
				}

				g.BeeTurn()
				time.Sleep(500 * time.Millisecond) // Slight delay between turns
			}
		}
	}()
}
