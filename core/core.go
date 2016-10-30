package core

import (
	"fmt"
	"time"
)

type Level string
type State string

type Event struct {
	level Level
	state State
}

func NewEvent(level Level, state State) (Event) {
	return Event{level:level, state:state}
}

// Responsible for checking outside data sources
type Monitor interface {
	Status() (Event)
	Start()
}

// Responsible fo triggering events
type Trigger interface {
	Fire()
	Start()
}

func GoMalina(monitors []Monitor, triggers []Trigger) {
	for idx, element := range monitors {
		fmt.Printf("Starting monitor#%d: %s\n", idx, element)
		go element.Start()
	}

	for idx, element := range triggers {
		fmt.Printf("Starting trigger#%d: %s\n", idx, element)
		go element.Start()
	}

	ticker := time.NewTicker(5 * time.Second)
	for tick := range ticker.C {
		tick.Clock()
		for _, element := range monitors {
			fmt.Printf("Status Check %s: \n", element.Status())
		}
	}
}