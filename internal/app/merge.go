package app

import (
	"auto-deploy-staging/internal/domain"
	"auto-deploy-staging/internal/parse"
	"encoding/json"
	"fmt"
	"os"
)

func Merge() []domain.Service {
	dataConfig, err := os.ReadFile("config.json")
	if err != nil {
		panic(fmt.Errorf("не удалось прочитать config.json: %w", err))
	}
	var config domain.Config
	if err := json.Unmarshal(dataConfig, &config); err != nil {
		panic(fmt.Errorf("ошибка разбора config.json: %w", err))
	}
	workspaceData, err := os.ReadFile(config.WorkspaceLocation)
	if err != nil {
		panic(fmt.Errorf("не удалось прочитать workspace: %w", err))
	}
	workspace, err := parse.Workspace(workspaceData)
	if err != nil {
		panic(fmt.Errorf("ошибка разбора workspace: %w", err))
	}
	_ = workspace

	dataSettingServices, err := os.ReadFile("services.json")
	if err != nil {
		panic(fmt.Errorf("не удалось прочитать service.json: %w", err))
	}
	var settingService []domain.SettingService
	if err := json.Unmarshal(dataSettingServices, &settingService); err != nil {
		panic(fmt.Errorf("ошибка разбора service.json: %w", err))
	}
	services := Convert(settingService, workspace)
	return services

}

func Convert(settings []domain.SettingService, packages map[string]domain.Package) []domain.Service {
	var result []domain.Service
	for _, setting := range settings {
		pkg, ok := packages[setting.WorkspaceName]
		if !ok {
			fmt.Printf("Предупреждение: пакет для workspace %q не найден\n", setting.WorkspaceName)
			continue
		}

		service := domain.Service{
			Name:       setting.WorkspaceName,
			Repository: pkg.Repository,
			Disabled:   setting.Disabled,
		}
		result = append(result, service)
	}
	return result
}
