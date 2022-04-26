package main

type State interface {
	DoClock(c Context, hour int)
	DoUse(c Context)
	DoAlarm(c Context)
	DoPhone(c Context)
}
