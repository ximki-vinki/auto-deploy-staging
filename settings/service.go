package settings

type Service struct {
	Name           string `json:"name"`
	WorkspaceName  string `json:"workspace_name"`
	KubernetesName string `json:"Kubernetes_name"`
	Disabled       bool   `json:"disabled"`
}
