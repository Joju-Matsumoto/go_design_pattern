package main

type Enemy int

const (
	Zombie Enemy = iota
	Skeleton
	Creeper
)

func (e Enemy) String() string {
	switch e {
	case Zombie:
		return "Zombie"
	case Skeleton:
		return "Skeleton"
	case Creeper:
		return "Creeper"
	}
	return "<INVALID>"
}

var Enemies = []Enemy{Zombie, Skeleton, Creeper}
