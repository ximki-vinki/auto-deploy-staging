package app

import (
	"auto-deploy-staging/internal/domain"
	"auto-deploy-staging/internal/parse"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const gitLabURL = "https://git.greensight.ru/api/v4/projects/auchan%2Fdevops%2Fdeveloper-pipelines/repository/files/README.md/raw?ref=master"
const token = "glpat-yGP2XiqqJY5JGHjD6Xz62m86MQp1OjRlCA.01.0y02iiljw"

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

func fetchFromGitLab() ([]byte, error) {
	url := os.Getenv("GITLAB_GET_README")
	token := os.Getenv("GITLAB_PRIVATE_TOKEN")
	if url == "" {
		return nil, fmt.Errorf("GITLAB_GET_README is not set")
	}
	if token == "" {
		return nil, fmt.Errorf("GITLAB_PRIVATE_TOKEN is not set")
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("PRIVATE-TOKEN", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got HTTP %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
