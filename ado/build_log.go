package ado

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c Client) GetBuildLog(buildId, logId int) (*string, error) {
	endpoint := fmt.Sprintf("%s/%s/%s/_apis/build/builds/%d/logs/%d?$format=text&api-version=5.0-preview.2", c.baseUrl, c.organization, c.project, buildId, logId)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.makeHttpRequest(req)
	if err != nil {
		return nil, fmt.Errorf("Error making HTTP Request: %s", err)
	}

	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error parsing HTTP Response: %s", err)
	}

	body := string(responseData)
	return &body, nil
}
