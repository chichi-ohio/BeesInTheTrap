package game

import "fmt"

type Player struct {
	HitPoints int
	MaxHP     int
	HitCount  int
}

func NewPlayer() *Player {
	return &Player{
		HitPoints: 100,
		MaxHP:     100,
		HitCount:  0,
	}
}

func (p *Player) IsAlive() bool {
	return p.HitPoints > 0
}

func (p *Player) TakeDamage(damage int) {
	p.HitPoints -= damage
	if p.HitPoints < 0 {
		p.HitPoints = 0
	}
}

func (p *Player) Hit() {
	p.HitCount++
}

func (p *Player) Status() string {
	return fmt.Sprintf("Player Status: %d/%d HP, Hits: %d", p.HitPoints, p.MaxHP, p.HitCount)
}
