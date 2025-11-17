package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParsePackages_OnePackage(t *testing.T) {
	yamlData := []byte(`
pim-pim:
  repository: git@git.test.ru:auchan/pim/pim/pim.git
`)

	packages, _ := ParsePackages(yamlData)

	require.Len(t, packages, 1)

	pkg := packages[0]
	assert.Equal(t, "pim-pim", pkg.Name)
	assert.Equal(t, "git@git.test.ru:auchan/pim/pim/pim.git", pkg.Repository)
}
