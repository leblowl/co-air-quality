package test

import (
	"co-air-quality.api/src/app"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Request(a app.App, req *http.Request) *httptest.ResponseRecorder {
	var rr = httptest.NewRecorder()

  a.Handler.ServeHTTP(rr, req)

	return rr
}

func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d", expected, actual)
	}
}

