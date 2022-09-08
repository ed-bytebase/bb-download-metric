package main

import (
	"context"
	"fmt"

	"github.com/bb-download-metric/metric"
)

// MetricReporter is the API message for metric reporter.
type MetricReporter struct {
	reporter   metric.Reporter
	collectors map[string]metric.Collector
}

// Run will collect the metric and send to the reporter.
func (m *MetricReporter) Run() error {
	ctx := context.Background()

	for name, collector := range m.collectors {
		fmt.Printf("run metric collector: %s\n", name)
		metricList, err := collector.Collect(ctx)

		if err != nil {
			fmt.Printf("failed to collect metric: %s with error %s\n", name, err.Error())
			continue
		}

		for _, metric := range metricList {
			fmt.Println(metric)
			m.report(metric)
		}
	}
	return nil
}

// Register will register a metric collector.
func (m *MetricReporter) Register(metricName metric.Name, collector metric.Collector) {
	m.collectors[string(metricName)] = collector
}

func (m *MetricReporter) report(metric *metric.Metric) {
	if err := m.reporter.Report(metric); err != nil {
		fmt.Printf("failed to report metric %s", metric.Name)
	}
}
