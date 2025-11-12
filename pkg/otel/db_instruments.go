package otel

import (
	"errors"
	"fmt"

	"go.opentelemetry.io/otel/metric"
)

type DBStatsInstruments struct {
	connectionMaxOpen   metric.Int64ObservableGauge
	connectionOpen      metric.Int64ObservableGauge
	connectionWaitTotal metric.Int64ObservableCounter
}

func NewDBStatsInstruments(meter metric.Meter) (*DBStatsInstruments, error) {
	var instruments DBStatsInstruments
	var err, e error

	subsystem := "connection"

	if instruments.connectionOpen, e = meter.Int64ObservableGauge(subsystem); e != nil {
		e = fmt.Errorf("failed to create instruments connectionOpen, %w", e)
		err = errors.Join(err, e)
	}

	if instruments.connectionMaxOpen, e = meter.Int64ObservableGauge(subsystem); e != nil {
		e = fmt.Errorf("failed to create instruments connectionMaxOpne, %w", e)
		err = errors.Join(err, e)
	}

	if instruments.connectionWaitTotal, e = meter.Int64ObservableCounter(subsystem); e != nil {
		e = fmt.Errorf("failed to create instruments connectionWaitTotal, %w", e)
		err = errors.Join(err, e)
	}

	return &instruments, err
}
