package otel

import (
	"fmt"
	"go.opentelemetry.io/otel/metric"
	"strings"
)

type dbStatsInstruments struct {
	connectionMaxOpen   metric.Int64ObservableGauge
	connectionOpen      metric.Int64ObservableGauge
	connectionWaitTotal metric.Int64ObservableCounter
}

type instruments struct {
	latency metric.Float64Histogram
}

const namespace = ""

func newInstruments(meter metric.Meter) (*instruments, error) {
	var instruments instruments
	var err error
	if instruments.latency, err = meter.Float64Histogram(
		strings.Join([]string{namespace, "latency"}, "."),
		metric.WithDescription("the latency of calls in milliseconds"),
		metric.WithUnit("ms"),
	); err != nil {
		return nil, fmt.Errorf("failed to create latency instrument, %v", err)
	}

	return &instruments, nil
}

func newDBStatsInstruments(meter metric.Meter) (*dbStatsInstruments, error) {
	var instruments dbStatsInstruments
	var err error

	subsystem := "connection"

	if instruments.connectionOpen, err = meter.Int64ObservableGauge(subsystem); err != nil {
		return nil, fmt.Errorf("failed to create instruments connectionOpen")
	}

	if instruments.connectionMaxOpen, err = meter.Int64ObservableGauge(subsystem); err != nil {
		return nil, fmt.Errorf("failed to create instruments connectionMaxOpne")
	}

	if instruments.connectionWaitTotal, err = meter.Int64ObservableCounter(subsystem); err != nil {
		return nil, fmt.Errorf("failed to create instruments connection")
	}

	return &instruments, nil
}
