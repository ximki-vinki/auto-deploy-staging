package example

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Pipeline struct {
	ID     int    `json:"id"`
	WebURL string `json:"web_url"`
}

func triggerPipeline(client *http.Client, baseURL, token, projectID, ref string, vars map[string]string) (*Pipeline, error) {
	apiURL := fmt.Sprintf("%s/projects/%s/trigger/pipeline", baseURL, url.PathEscape(projectID))

	payload := map[string]string{
		"ref": ref,
	}

	// Добавляем переменные CI
	for k, v := range vars {
		payload[fmt.Sprintf("variables[%s]", k)] = v
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(string(jsonData)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("PRIVATE-TOKEN", token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("failed to trigger pipeline: %s", resp.Status)
	}

	var result Pipeline
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
