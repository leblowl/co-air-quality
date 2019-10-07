package test

import (
	"co-air-quality.api/src/precursor/model"
	"testing"
)

func TestGetObservations(t *testing.T) {
	obs, err = model.GetObservations()

	if err != nil {
		t.Fatal(err)
	}
}
