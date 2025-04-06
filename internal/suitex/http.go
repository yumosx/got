package suitex

import (
	"bytes"
	"net/http"
	"net/http/httptest"
)

// MockPostResponse Post(gin.Server, "/user/", json)
func MockPostResponse(server http.Handler, path string, data []byte) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	server.ServeHTTP(resp, req)
	return resp, err
}
