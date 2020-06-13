package fetch

import "testing"

func TestMetrics_StartedAtGetter(t *testing.T) {
	metrics := &Metrics{}
	if metrics.StartedAt() != metrics.startedAt {
		t.Fail()
	}
}

func TestMetrics_StartedAtGetter_ShouldBeZero(t *testing.T) {
	metrics := &Metrics{}
	if !metrics.StartedAt().IsZero() {
		t.Fail()
	}
}

func TestMetrics_CompletedAtGetter(t *testing.T) {
	metrics := &Metrics{}
	if metrics.CompletedAt() != metrics.completedAt {
		t.Fail()
	}
}

func TestMetrics_CompletedAtGetter_ShouldBeZero(t *testing.T) {
	metrics := &Metrics{}
	if !metrics.CompletedAt().IsZero() {
		t.Fail()
	}
}

func TestMetrics_DurationGetter(t *testing.T) {
	metrics := &Metrics{}
	if metrics.Duration() != metrics.duration {
		t.Fail()
	}
}

func TestMetrics_DurationGetter_ShouldBeZero(t *testing.T) {
	metrics := &Metrics{}
	if metrics.Duration() != 0 {
		t.Fail()
	}
}

func TestMetrics_Start_StartedAt_ShouldnBeZero(t *testing.T) {
	metrics := &Metrics{}
	metrics.Start()

	if metrics.StartedAt().IsZero() {
		t.Fail()
	}
}

func TestMetrics_Complete_CompletedAt_ShouldnBeZero(t *testing.T) {
	metrics := &Metrics{}
	metrics.Start()
	metrics.Complete()

	if metrics.CompletedAt().IsZero() {
		t.Fail()
	}
}

func TestMetrics_Complete_Duration_ShouldBeCalculated(t *testing.T) {
	metrics := &Metrics{}
	metrics.Start()
	metrics.Complete()

	if metrics.Duration() != metrics.CompletedAt().Sub(metrics.StartedAt()) {
		t.Fail()
	}
}
