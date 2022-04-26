package main

import "fmt"

func main() {
	agent := NewAgentJJ()

	for turn := 0; turn < 100; turn++ {
		agent.Update(turn)
		fmt.Println(agent)
	}
}
