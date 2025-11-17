package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseWorkspace_OnePackage(t *testing.T) {
	yamlData := []byte(`
services:
  pim-pim:
    repository: git@git.test.ru:auchan/pim/pim/pim.git
`)

	packages, _ := Workspace(yamlData)

	require.Len(t, packages, 1)

	pkg := packages["pim-pim"]
	assert.Equal(t, "git@git.test.ru:auchan/pim/pim/pim.git", pkg.Repository)
}
