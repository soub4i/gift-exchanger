package v1_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/soub4i/giftsxchanger/pkg/api"
	"github.com/stretchr/testify/assert"
)

func Test_Health(t *testing.T) {
	router := api.Register("v1")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "healthy")
}
