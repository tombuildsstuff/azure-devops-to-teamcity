package ado

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type getBuildsResponse struct {
	Builds []Build `json:"value"`
}

type Build struct {
	Id          int    `json:"id"`
	BuildNumber string `json:"buildNumber"`
	Status      string `json:"status"`
	Result      string `json:"result"`
}

func (c Client) GetAvailableBuilds() (*[]Build, error) {
	endpoint := fmt.Sprintf("%s/%s/%s/_apis/build/builds?api-version=5.0-preview.5", c.baseUrl, c.organization, c.project)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.makeHttpRequest(req)
	if err != nil {
		return nil, fmt.Errorf("Error making HTTP Request: %s", err)
	}

	var response *getBuildsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("Error deserializing JSON: %s", err)
	}

	return &response.Builds, nil
}
