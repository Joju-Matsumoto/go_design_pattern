package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	bell := NewBellPrinter()

	// commands := []string{"USE", "ALARM", "PHONE"}
	commands := []Command{UseCommand}

	go func() {
		for {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(3000)))
			bell.Use(commands[rand.Intn(len(commands))])
		}
	}()

	for {
		for now := 0; now < 24; now++ {
			bell.setClock(now)
			time.Sleep(1 * time.Second)
		}
	}
}
