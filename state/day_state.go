package main

import "fmt"

var ds *DayState

func init() {
	ds = &DayState{}
}

func GetDayState() State {
	return ds
}

type DayState struct {
}

func (ds DayState) String() string {
	return fmt.Sprintf("DayState")
}

func (ds *DayState) DoClock(c Context, hour int) {
	if hour < 9 || 17 <= hour {
		c.changeState(GetNightState())
	}
}

func (ds *DayState) DoUse(c Context) {
	c.recording("Day time Use")
}
func (ds *DayState) DoAlarm(c Context) {
	c.callSecurityCenter("Emergency Bell (Day)")
}

func (ds *DayState) DoPhone(c Context) {
	c.recording("Day time Phone")
}
