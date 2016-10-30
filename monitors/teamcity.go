package monitors

import (
	"time"
	"github.com/bbucko/symmetrical-happiness/core"
	"github.com/bbucko/symmetrical-happiness/core/level"
	"github.com/bbucko/symmetrical-happiness/core/state"
)

type TeamCity struct {
	url       string
	projectId string

	status    string
	ok        bool
}

type option func(*TeamCity)

func (f *TeamCity) Option(opts ...option) {
	for _, opt := range opts {
		opt(f)
	}
}

func (tc TeamCity) String() string {
	return "teamcity: " + tc.url
}

func (tc TeamCity) Start() {
	ticker := time.NewTicker(5 * time.Second)
	for tick := range ticker.C {
		tick.Clock()
		tc.check()
	}
}

func (tc TeamCity) check() {
	//send request
	//parse response
	//set status
	tc.ok = true
	tc.status = "failing"
}

func (TeamCity) Status() (core.Event) {
	return core.NewEvent(level.High, state.OK)
}

func NewTeamCity(opts ...option) (*TeamCity) {
	tc := TeamCity{}
	tc.Option(opts... )
	return &tc
}

//Options
func URL(url string) option {
	return func(tc *TeamCity) {
		tc.url = url
	}
}

func ProjectID(projectId string) option {
	return func(tc *TeamCity) {
		tc.projectId = projectId
	}
}