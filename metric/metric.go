// Package metric is the interfaces for telemetry metrics.
package metric

// Name is the metric name.
type Name string

// Metric is the API message for metric.
type Metric struct {
	Name   Name
	Value  int
	Labels map[string]string
}
