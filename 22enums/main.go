package main

import "fmt"

type severState int

const (
	START severState = iota
	RUNNING
	STOPPED
)

var stateName = map[severState]string{
	START:   "start",
	RUNNING: "running",
	STOPPED: "stopped",
}

func (ss severState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(START)
	fmt.Println(ns)
	ns = transition(RUNNING)
	fmt.Println(ns)
	ns = transition(STOPPED)
	fmt.Println(ns)
}

func transition(s severState) severState {
	switch s {
	case START:
		return RUNNING
	case RUNNING:
		return STOPPED
	case STOPPED:
		return START
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
