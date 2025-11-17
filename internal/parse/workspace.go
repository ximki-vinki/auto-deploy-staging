package parse

import (
	"auto-deploy-staging/internal/domain"

	"gopkg.in/yaml.v3"
)

type workspaceConfig struct {
	Services map[string]domain.Package `yaml:"services"`
}

func Workspace(yamlData []byte) (map[string]domain.Package, error) {
	var workspace workspaceConfig
	err := yaml.Unmarshal(yamlData, &workspace)
	if err != nil {
		return nil, err
	}
	return workspace.Services, nil
}
