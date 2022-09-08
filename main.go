package main

import (
	"fmt"
	"log"

	"github.com/bb-download-metric/collector"
	"github.com/bb-download-metric/docker"
	"github.com/bb-download-metric/github"
	"github.com/bb-download-metric/metric"
	"github.com/bb-download-metric/metric/segment"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("cron job started")

	loadEnv("prod")

	r := segment.NewReporter(viper.GetString("SEGMENT_KEY"), "cron")
	reporter := &MetricReporter{
		reporter:   r,
		collectors: make(map[string]metric.Collector),
	}
	reporter.Register(collector.AssetsDownloadCountMetricName, &collector.GithubCollector{
		Repository:    "bytebase/bytebase",
		GithubService: github.NewService(viper.GetString("GITHUB_TOKEN")),
	})
	reporter.Register(collector.DockerImagePullCountMetricName, &collector.DockerCollector{
		Repository:    "bytebase/bytebase",
		DockerService: docker.NewService(viper.GetString("DOCKER_TOKEN")),
	})

	if err := reporter.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("cron job finished")
}

func loadEnv(env string) error {
	viper.SetConfigName(env)
	viper.SetConfigType("env")
	viper.AddConfigPath("./env")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("an error occurred reading the config file: %s ", err)
	}

	return nil
}
