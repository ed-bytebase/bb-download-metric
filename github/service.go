package github

import (
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

// Service is the GitHub Service.
type Service struct {
	token    string
	endpoint string
}

// NewService returns the instance of GitHub service.
func NewService(token string) *Service {
	return &Service{
		token:    token,
		endpoint: "https://api.github.com",
	}
}

// ListRelease lists the release for the repo.
// API docs: https://docs.github.com/en/rest/releases/releases#list-releases
func (s *Service) ListRelease(repository string) ([]*Release, error) {
	endpoint := fmt.Sprintf("%s/repos/%s/releases", s.endpoint, repository)
	resp, body, errs := gorequest.New().Get(endpoint).
		Set("Authorization", fmt.Sprintf("Bearer %s", s.token)).
		Set("Accept", "application/vnd.github+json").
		End()

	if len(errs) > 0 {
		return nil, fmt.Errorf("list pull request files failed with error %s", errs[0].Error())
	}

	if err := checkResponse(resp.StatusCode, body); err != nil {
		return nil, err
	}

	res := new([]*Release)
	if err := json.Unmarshal([]byte(body), res); err != nil {
		return nil, err
	}

	return *res, nil
}

func checkResponse(code int, body string) error {
	if code >= 200 && code < 300 {
		return nil
	}

	return fmt.Errorf("request failed with response code %d and body %s", code, body)
}
