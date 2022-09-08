package collector

import (
	"context"

	"github.com/bb-download-metric/github"
	"github.com/bb-download-metric/metric"
)

// GithubCollector is the collector for GitHub.
type GithubCollector struct {
	Repository    string
	GithubService *github.Service
}

// Collect returns the list of metric.
func (c *GithubCollector) Collect(ctx context.Context) ([]*metric.Metric, error) {
	releaseList, err := c.GithubService.ListRelease(c.Repository)
	if err != nil {
		return nil, err
	}

	var res []*metric.Metric

	for _, release := range releaseList {
		if release.Draft || release.Prerelease {
			continue
		}

		total := 0
		for _, asset := range release.Assets {
			total += asset.DownloadCount
		}

		res = append(res, &metric.Metric{
			Name:  AssetsDownloadCountMetricName,
			Value: total,
			Labels: map[string]string{
				"platform":   "GitHub",
				"repository": c.Repository,
				"tag":        release.TagName,
			},
		})
	}

	return res, nil
}
