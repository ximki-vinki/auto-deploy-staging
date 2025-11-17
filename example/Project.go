package example

import (
	"encoding/json"
	"os"
)

type Project struct {
	ID     string `json:"id"`
	Branch string `json:"branch,omitempty"`
}

func readProjectsFromFile(path string) ([]Project, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var projects []Project
	err = json.Unmarshal(data, &projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
