package collector

import "github.com/bb-download-metric/metric"

const (
	// AssetsDownloadCountMetricName is the metric name for asset download count.
	AssetsDownloadCountMetricName metric.Name = "bb.download.github"

	// DockerImagePullCountMetricName is the metric name for docker image pull count.
	DockerImagePullCountMetricName metric.Name = "bb.download.docker"
)
