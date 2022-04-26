package main

import (
	"math/rand"
)

type Wizard struct {
}

var wizard *Wizard

func init() {
	wizard = &Wizard{}
}

func getWizard() Mode {
	return wizard
}

func (w Wizard) String() string {
	return "\033[33mWizard\033[0m"
}

func (w *Wizard) Update(h Hero) {
	if len(h.getEnemies()) < 5 {
		h.setMode(getAttacker())
	}
}

func (w *Wizard) Act(h Hero) {
	n := rand.Intn(5)
	for i := 0; i < n; i++ {
		h.kill()
	}
}
