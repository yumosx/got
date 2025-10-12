package mockhttp

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClientConfig(t *testing.T) {
	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {})
	config := NewClientConfig(
		WithPath("/path"),
		WithMethod("POST"),
		WithData([]byte(`{"key": "value"}`)),
		WithHeaders(map[string]string{"Content-Type": "application/json"}),
		WithServer(handler),
	)

	assert.Equal(t, config.path, "/path")
	assert.Equal(t, config.method, "POST")
	assert.Equal(t, config.data, []byte(`{"key": "value"}`))
	assert.Equal(t, config.headers, map[string]string{"Content-Type": "application/json"})
	assert.NotEmpty(t, config.server)
}

func TestPost(t *testing.T) {
	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte(`{"status":"ok"}`))
	})

	config := NewClientConfig(
		WithPath("/test"),
		WithServer(handler))

	resp, err := Post(config)
	require.NoError(t, err)
	assert.Equal(t, resp.Code, http.StatusOK)
	assert.Equal(t, resp.Body.String(), `{"status":"ok"}`)
}
