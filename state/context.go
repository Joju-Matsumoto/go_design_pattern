package main

type Context interface {
	setClock(hour int)
	changeState(s State)
	callSecurityCenter(msg string)
	recording(msg string)
}
