package docker

// Repository is the API message for a docker image.
type Repository struct {
	User           string `json:"user"`
	Name           string `json:"name"`
	Namespace      string `json:"namespace"`
	RepositoryType string `json:"repository_type"`
	IsPrivate      bool   `json:"is_private"`
	PullCount      int    `json:"pull_count"`
}
