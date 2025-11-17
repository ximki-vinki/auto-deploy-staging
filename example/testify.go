package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadProjectsFromFile(t *testing.T) {
	projects, err := readProjectsFromFile("test_projects.json")

	require.NoError(t, err)
	assert.Len(t, projects, 2)
	assert.Equal(t, "my-group/my-project-1", projects[0].ID)
	assert.Equal(t, "develop", projects[0].Branch)
}
