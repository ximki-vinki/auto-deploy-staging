package domain

type SettingService struct {
	WorkspaceName  string `json:"workspace_name"`
	KubernetesName string `json:"Kubernetes_name"`
	Disabled       bool   `json:"disabled"`
}
