package main

type Attacker struct {
}

var attacker *Attacker

func init() {
	attacker = &Attacker{}
}

func getAttacker() Mode {
	return attacker
}

func (a Attacker) String() string {
	return "\033[32mAttacker\033[0m"
}

// implements Mode

func (a *Attacker) Update(h Hero) {
	enemies := h.getEnemies()
	if len(enemies) > 5 {
		h.setMode(getWizard())
	}
}

func (a *Attacker) Act(h Hero) {
	h.kill()
}
