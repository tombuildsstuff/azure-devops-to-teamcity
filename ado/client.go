package ado

import (
	"net/http"
)

type Client struct {
	baseUrl      string
	organization string
	project      string
	token        string
}

func NewClient(baseUrl, organization, project, token string) Client {
	return Client{
		baseUrl:      baseUrl,
		organization: organization,
		project:      project,
		token:        token,
	}
}

func (c Client) makeHttpRequest(request *http.Request) (*http.Response, error) {
	request.SetBasicAuth("", c.token)
	return http.DefaultClient.Do(request)
}
