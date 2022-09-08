package collector

import (
	"context"

	"github.com/bb-download-metric/docker"
	"github.com/bb-download-metric/metric"
)

// DockerCollector is the collector for Docker.
type DockerCollector struct {
	Repository    string
	DockerService *docker.Service
}

// Collect returns the list of metric.
func (c *DockerCollector) Collect(ctx context.Context) ([]*metric.Metric, error) {
	repo, err := c.DockerService.GetRepository(c.Repository)
	if err != nil {
		return nil, err
	}

	return []*metric.Metric{
		{
			Name:  DockerImagePullCountMetricName,
			Value: repo.PullCount,
			Labels: map[string]string{
				"platform":   "Docker",
				"repository": c.Repository,
				"tag":        "latest",
			},
		},
	}, nil
}
