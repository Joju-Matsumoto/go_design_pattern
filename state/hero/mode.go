package main

type Mode interface {
	Update(h Hero)
	Act(h Hero)
}
