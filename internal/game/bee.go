package game

import (
	"fmt"
)

type BeeType int

const (
	Queen BeeType = iota
	Worker
	Drone
)

type Bee struct {
	Type      BeeType
	HitPoints int
	MaxHP     int
	Damage    int
	Name      string
}

func NewQueen() *Bee {
	return &Bee{
		Type:      Queen,
		HitPoints: 100,
		MaxHP:     100,
		Damage:    10,
		Name:      "Queen Bee",
	}
}

func NewWorker() *Bee {
	return &Bee{
		Type:      Worker,
		HitPoints: 75,
		MaxHP:     75,
		Damage:    5,
		Name:      "Worker Bee",
	}
}

func NewDrone() *Bee {
	return &Bee{
		Type:      Drone,
		HitPoints: 60,
		MaxHP:     60,
		Damage:    1,
		Name:      "Drone Bee",
	}
}

func (b *Bee) IsAlive() bool {
	return b.HitPoints > 0
}

func (b *Bee) TakeDamage() int {
	var damage int
	switch b.Type {
	case Queen:
		damage = 10
	case Worker:
		damage = 25
	case Drone:
		damage = 30
	}

	b.HitPoints -= damage
	if b.HitPoints < 0 {
		b.HitPoints = 0
	}
	return damage
}

func (b *Bee) AttackMessage(damage int) string {
	return fmt.Sprintf("Direct Hit! You took %d hit points from a %s!", damage, b.Name)
}

func (b *Bee) StingMessage() string {
	return fmt.Sprintf("Sting! You just got stung by a %s!", b.Name)
}

func (b *Bee) String() string {
	return fmt.Sprintf("%s: %d/%d HP", b.Name, b.HitPoints, b.MaxHP)
}
