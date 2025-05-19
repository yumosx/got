package mockhttp

import (
	"bytes"
	"net/http"
	"net/http/httptest"
)

type HTTPClientConfig struct {
	server  http.Handler
	path    string
	method  string
	data    []byte
	headers map[string]string
}

type HTTPConfigOption interface {
	Option(config *HTTPClientConfig)
}

type HTTPConfigOptionFunc func(config *HTTPClientConfig)

func (fn HTTPConfigOptionFunc) Option(config *HTTPClientConfig) {
	fn(config)
}

func WithHeaders(headers map[string]string) HTTPConfigOption {
	return HTTPConfigOptionFunc(func(config *HTTPClientConfig) {
		config.headers = headers
	})
}

func WithServer(server http.Handler) HTTPConfigOption {
	return HTTPConfigOptionFunc(func(config *HTTPClientConfig) {
		config.server = server
	})
}

func WithPath(path string) HTTPConfigOption {
	return HTTPConfigOptionFunc(func(config *HTTPClientConfig) {
		config.path = path
	})
}

func WithMethod(method string) HTTPConfigOption {
	return HTTPConfigOptionFunc(func(config *HTTPClientConfig) {
		config.method = method
	})
}

func WithData(data []byte) HTTPConfigOption {
	return HTTPConfigOptionFunc(func(config *HTTPClientConfig) {
		config.data = data
	})
}

func NewClientConfig(options ...HTTPConfigOption) *HTTPClientConfig {
	config := &HTTPClientConfig{}
	for _, opt := range options {
		opt.Option(config)
	}

	return config
}

func Post(config *HTTPClientConfig) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(http.MethodPost, config.path, bytes.NewBuffer(config.data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	config.server.ServeHTTP(resp, req)

	return resp, nil
}

func Get(config *HTTPClientConfig) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(http.MethodGet, config.path, bytes.NewBuffer(nil))
	if err != nil {
		return nil, err
	}
	resp := httptest.NewRecorder()
	config.server.ServeHTTP(resp, req)
	return resp, nil
}

func Do(config *HTTPClientConfig) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(config.method, config.path, bytes.NewBuffer(config.data))
	if err != nil {
		return nil, err
	}

	for key, value := range config.headers {
		req.Header.Set(key, value)
	}

	resp := httptest.NewRecorder()
	config.server.ServeHTTP(resp, req)
	return resp, nil
}
