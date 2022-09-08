package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bb-download-metric/collector"
	"github.com/bb-download-metric/docker"
	"github.com/bb-download-metric/github"
	"github.com/bb-download-metric/metric"
	"github.com/bb-download-metric/metric/segment"
)

func main() {
	fmt.Println("cron job started")

	client := segment.NewReporter(os.Getenv("SEGMENT_KEY"), "download_metric")
	defer client.Close()

	reporter := &MetricReporter{
		reporter:   client,
		collectors: make(map[string]metric.Collector),
	}
	reporter.Register(collector.AssetsDownloadCountMetricName, &collector.GithubCollector{
		Repository:    "bytebase/bytebase",
		GithubService: github.NewService(os.Getenv("GITHUB_TOKEN")),
	})
	reporter.Register(collector.DockerImagePullCountMetricName, &collector.DockerCollector{
		Repository:    "bytebase/bytebase",
		DockerService: docker.NewService(os.Getenv("DOCKER_TOKEN")),
	})

	if err := reporter.Run(); err != nil {
		log.Printf("reporter failed with error %s\n", err.Error())
	}

	fmt.Println("cron job finished")
}
