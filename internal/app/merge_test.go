package app

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {
	// Подготовка тестовых данных
	configContent := `{"token": "test", "workspace_location": "workspace.yaml"}`
	workspaceContent := `
services:
  service1:
    repository: repo1
  service2:
    repository: repo2
`
	servicesContent := `[{"workspace_name": "service1", "Kubernetes_name": "k8s1", "disabled": false}]`

	err := os.WriteFile("config.json", []byte(configContent), 0644)
	require.NoError(t, err)
	defer os.Remove("config.json")

	err = os.WriteFile("workspace.yaml", []byte(workspaceContent), 0644)
	require.NoError(t, err)
	defer os.Remove("workspace.yaml")

	err = os.WriteFile("services.json", []byte(servicesContent), 0644)
	require.NoError(t, err)
	defer os.Remove("services.json")

	result := Merge()

	assert.Len(t, result, 1)

	if len(result) > 0 {
		assert.Equal(t, "service1", result[0].Name)
		assert.Equal(t, "repo1", result[0].Repository)
	}
}
