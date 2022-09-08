package docker

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

// NewService returns the instance of Docker service.
func NewService(token string) *Service {
	return &Service{
		token:    token,
		endpoint: "https://hub.docker.com/v2",
	}
}

// GetRepository gets the repository.
func (s *Service) GetRepository(repository string) (*Repository, error) {
	endpoint := fmt.Sprintf("%s/repositories/%s", s.endpoint, repository)
	resp, body, errs := gorequest.New().Get(endpoint).
		Set("Authorization", fmt.Sprintf("Bearer %s", s.token)).
		Set("Accept", "application/json").
		End()

	if len(errs) > 0 {
		return nil, fmt.Errorf("get repository failed with error %s", errs[0].Error())
	}

	if err := checkResponse(resp.StatusCode, body); err != nil {
		return nil, err
	}

	res := new(Repository)
	if err := json.Unmarshal([]byte(body), res); err != nil {
		return nil, err
	}

	return res, nil
}

func checkResponse(code int, body string) error {
	if code >= 200 && code < 300 {
		return nil
	}

	return fmt.Errorf("request failed with response code %d and body %s", code, body)
}
