package v1_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/soub4i/giftsxchanger/pkg/api"
	"github.com/soub4i/giftsxchanger/pkg/datastore"
	"github.com/stretchr/testify/assert"
)

func Test_Shuffle(t *testing.T) {

	version := "v1"
	ds := datastore.GetDS()
	router := api.Register(version)
	ds.Seed()

	assert.Empty(t, ds.GetExchanges())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/gift-exchange", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, ds.GetExchanges())

}

func Test_Shuffle_Failure(t *testing.T) {
	version := "v1"
	router := api.Register(version)
	w := httptest.NewRecorder()

	for i := 0; i < 15; i++ {
		req, _ := http.NewRequest("POST", "/v1/gift-exchange", nil)
		router.ServeHTTP(w, req)

	}
	r := w.Result()
	assert.Equal(t, 500, r.StatusCode)

}
