package parse

import (
	"auto-deploy-staging/internal/domain"

	"gopkg.in/yaml.v3"
)

func Workspace(yamlData []byte) (map[string]domain.Package, error) {
	var packages map[string]domain.Package
	err := yaml.Unmarshal(yamlData, &packages)
	if err != nil {
		return nil, err
	}

	return packages, nil
}
