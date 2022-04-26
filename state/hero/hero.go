package main

import (
	"fmt"
	"math/rand"
)

type Hero interface {
	Update(turn int)
	getTurn() int
	getEnemies() []Enemy
	setMode(m Mode)
	kill()
}

type AgentJJ struct {
	turn    int
	enemies []Enemy
	mode    Mode
}

func (a AgentJJ) String() string {
	return fmt.Sprintf("AgentJJ{turn: %v, mode: %v, enemies: %v}", a.turn, a.mode, a.enemies)
}

func NewAgentJJ() *AgentJJ {
	return &AgentJJ{
		turn: 0,
		mode: getAttacker(),
	}
}

func (a *AgentJJ) Update(turn int) {
	a.turn = turn

	n := rand.Intn(5)
	for i := 0; i < n; i++ {
		a.enemies = append(a.enemies, Enemies[rand.Intn(len(Enemies))])
	}

	a.mode.Update(a)
	a.mode.Act(a)
}

// implements Hero

func (a *AgentJJ) getTurn() int {
	return a.turn
}

func (a *AgentJJ) getEnemies() []Enemy {
	return a.enemies
}

func (a *AgentJJ) setMode(m Mode) {
	a.mode = m
}

func (a *AgentJJ) kill() {
	if len(a.enemies) > 0 {
		fmt.Printf("mode(%v): kill %v\n", a.mode, a.enemies[0])
		a.enemies = a.enemies[1:]
	}
}
