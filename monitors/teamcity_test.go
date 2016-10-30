package monitors

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestOptionURL(t *testing.T) {
	tc := TeamCity{}
	tc.Option(URL("abc"))
	assert.Equal(t, "abc", tc.url)
}

func TestProjectID(t *testing.T) {
	tc := TeamCity{}
	tc.Option(ProjectID("abc"))
	assert.Equal(t, "abc", tc.projectId)
}