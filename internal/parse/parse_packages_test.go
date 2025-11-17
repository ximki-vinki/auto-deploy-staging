package parse

import (
	"auto-deploy-staging/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePackages_ReturnsCorrectType(t *testing.T) {
	packages, _ := ParsePackages(nil)

	assert.Equal(t, []domain.Package{}, packages)
}
