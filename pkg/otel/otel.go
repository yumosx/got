package otel

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type Config struct {
	TracerProvider trace.TracerProvider
	Tracer         trace.Tracer
	MeterProvider  metric.MeterProvider
	Meter          metric.Meter
	Attributes     []attribute.KeyValue
	Version        string
}

type Option interface {
	Apply(config *Config)
}

type OptionFunc func(config *Config)

func (f OptionFunc) Apply(config *Config) {
	f(config)
}

func WithTracerProvider(tp trace.TracerProvider) Option {
	return OptionFunc(func(config *Config) {
		config.TracerProvider = tp
	})
}

func WithMeterProvider(meter metric.MeterProvider) Option {
	return OptionFunc(func(config *Config) {
		config.MeterProvider = meter
	})
}

func Version(name string) string {
	if name == "" {
		return "0.0.0"
	}
	return name
}

func NewConfig(name string, options ...Option) Config {
	config := Config{
		TracerProvider: otel.GetTracerProvider(),
		MeterProvider:  otel.GetMeterProvider(),
	}

	for _, opt := range options {
		opt.Apply(&config)
	}
	config.Tracer = config.TracerProvider.Tracer(
		name, trace.WithInstrumentationVersion(Version(config.Version)))
	config.Meter = config.MeterProvider.Meter(
		name, metric.WithInstrumentationVersion(Version(config.Version)))
	return config
}
