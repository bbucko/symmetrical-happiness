package main

import (
	"github.com/bbucko/symmetrical-happiness/core"
	"github.com/bbucko/symmetrical-happiness/monitors"
)

func main() {
	monitorsArr := []core.Monitor{
		monitors.NewTeamCity(monitors.URL("http://teamciy.com"), monitors.ProjectID("ProjectID")),
	}

	triggersArr := []core.Trigger{}

	core.GoMalina(monitorsArr, triggersArr)
}