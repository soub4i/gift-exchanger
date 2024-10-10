package v1_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/soub4i/giftsxchanger/pkg/api"
	"github.com/soub4i/giftsxchanger/pkg/datastore"
	"github.com/stretchr/testify/assert"
)

func Test_GetMembers(t *testing.T) {
	version := "v1"
	ds := datastore.GetDS()
	ds.Seed()
	router := api.Register(version)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v1/members", nil)
	router.ServeHTTP(w, req)
	var res []datastore.Member

	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Error(err)
	}
	t.Log(res)
	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, res)
}
func Test_GetMember(t *testing.T) {
	version := "v1"
	ds := datastore.GetDS()
	ds.Seed()
	router := api.Register(version)

	var tests = []struct {
		name    string
		path    string
		value   *datastore.Member
		wantErr bool
	}{
		{"I should get member", "/v1/members/1", &ds.Members[0], false},
		{"I should get 404 error", "/v1/members/90", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tt.path, nil)
			router.ServeHTTP(w, req)
			var res datastore.Member

			if tt.wantErr {
				assert.NotEqual(t, 200, w.Code)
			} else {
				if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
					t.Error(err)
				}
				assert.Equal(t, 200, w.Code)
				assert.Equal(t, res.Name, tt.value.Name)
			}

		})
	}

}

func Test_DeleteMember(t *testing.T) {
	version := "v1"
	ds := datastore.GetDS()
	router := api.Register(version)

	var tests = []struct {
		name    string
		path    string
		wantErr bool
	}{
		{"I should delete member", "/v1/members/1", false},
		{"I should get 404 error", "/v1/members/90", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds.Seed()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", tt.path, nil)
			router.ServeHTTP(w, req)

			if tt.wantErr {
				assert.NotEqual(t, 204, w.Code)
			} else {

				assert.Equal(t, 204, w.Code)
				assert.Nil(t, ds.GetMember("1"))
			}

		})
	}

}
func Test_AddMember(t *testing.T) {
	version := "v1"
	ds := datastore.GetDS()
	router := api.Register(version)

	var tests = []struct {
		name    string
		path    string
		payload []byte
		wantErr bool
	}{
		{"I should add member", "/v1/members", []byte(`{"id":"","name":"Abdel"}`), false},
		{"I should get  error", "/v1/members", []byte(""), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds.Seed()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", tt.path, bytes.NewReader(tt.payload))
			router.ServeHTTP(w, req)

			if tt.wantErr {
				assert.NotEqual(t, 201, w.Code)
			} else {

				assert.Equal(t, 201, w.Code)
				assert.Equal(t, ds.GetMember(strconv.Itoa(len(ds.Members))).Name, "Abdel")
			}

		})
	}

}

func Test_UpdateMember(t *testing.T) {
	version := "v1"
	ds := datastore.GetDS()
	router := api.Register(version)

	var tests = []struct {
		name    string
		path    string
		payload []byte
		wantErr bool
	}{
		{"I should update member", "/v1/members/1", []byte(`{"name":"Abdel"}`), false},
		{"I should get  error", "/v1/members/foo", []byte(""), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds.Seed()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PUT", tt.path, bytes.NewReader(tt.payload))
			router.ServeHTTP(w, req)

			if tt.wantErr {
				assert.NotEqual(t, 200, w.Code)
			} else {

				assert.Equal(t, 200, w.Code)
				assert.Equal(t, ds.GetMember("1").Name, "Abdel")
			}

		})
	}

}
