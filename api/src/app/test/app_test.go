package test

import (
	app "co-air-quality.api/src/app"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"reflect"
	"testing"
)

var a app.App

func TestMain(m *testing.M) {
	a = app.App{}

	var code int
	var err = a.Init(map[string]string{})

	if err == nil {
		code = m.Run()
	} else {
		code = 1
		// TODO: Figure out why logging doesn't appear to work in TestMain.
		// Yet, logging in app.Init works fine.
		log.Fatal("Fatal error during app initialization. Tests not run.")
	}
	os.Exit(code)
}

func TestIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := Request(a, req)

	CheckResponseCode(t, http.StatusOK, res.Code)

	var bodyExpected = map[string]bool{"alive": true}
	var contentTypeExpected = []string{"application/json"}

	var body map[string]bool
	var bodyJsonErr = json.Unmarshal(res.Body.Bytes(), &body)
	var headers = res.HeaderMap
	var contentType = headers["Content-Type"]

	if !reflect.DeepEqual(contentType, contentTypeExpected) {
		t.Errorf("Expected content type %v. Got %v",
			contentTypeExpected, contentType);
	}

	if bodyJsonErr != nil {
		t.Errorf("Error decoding JSON body %s", res.Body.String())
	}

	if !reflect.DeepEqual(body, bodyExpected) {
		t.Errorf("Expected %v. Got %v", bodyExpected, body)
	}
}
