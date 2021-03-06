package main

import "fmt"

type BellPrinter struct {
	state State
}

func NewBellPrinter() *BellPrinter {
	return &BellPrinter{state: GetNightState()}
}

func (b *BellPrinter) Use(c Command) {
	switch c {
	case UseCommand:
		b.state.DoUse(b)
	case AlarmCommand:
		b.state.DoAlarm(b)
	case PhoneCommand:
		b.state.DoPhone(b)
	}
}

func (b *BellPrinter) setClock(hour int) {
	fmt.Printf("time: %2d:00\n", hour)
	b.state.DoClock(b, hour)
}

func (b *BellPrinter) changeState(s State) {
	b.state = s
}

func (b *BellPrinter) callSecurityCenter(msg string) {
	fmt.Println("call:", msg)
}

func (b *BellPrinter) recording(msg string) {
	fmt.Println("record:", msg)
}
