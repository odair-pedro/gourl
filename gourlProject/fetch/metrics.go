package fetch

import (
	"time"
)

type Metrics struct {
	startedAt   time.Time
	completedAt time.Time
	duration    time.Duration
}

func (metrics *Metrics) StartedAt() time.Time {
	return metrics.startedAt
}

func (metrics *Metrics) CompletedAt() time.Time {
	return metrics.completedAt
}

func (metrics *Metrics) Duration() time.Duration {
	return metrics.duration
}

func (metrics *Metrics) Start() {
	metrics.startedAt = time.Now()
}

func (metrics *Metrics) Complete() {
	if !metrics.startedAt.IsZero() {
		metrics.completedAt = time.Now()
		metrics.duration = metrics.completedAt.Sub(metrics.startedAt)
	}
}
