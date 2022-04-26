package main

import (
	"fmt"
)

var ns *NightState

func init() {
	ns = &NightState{}
}

func GetNightState() State {
	return ns
}

type NightState struct {
}

func (ns NightState) String() string {
	return fmt.Sprintf("NightState{}")
}

// Stateを実装

func (ns *NightState) DoClock(c Context, hour int) {
	if 9 <= hour && hour < 17 {
		c.changeState(GetDayState())
	}
}

func (ns *NightState) DoUse(c Context) {
	c.callSecurityCenter("Emergency: Night time use")
}

func (ns *NightState) DoAlarm(c Context) {
	c.callSecurityCenter("Emergency Bell (Night)")
}

func (ns *NightState) DoPhone(c Context) {
	c.recording("Night time Phone")
}
