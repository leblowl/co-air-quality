package model

import (
	"errors"
	"time"
)

type PrecursorObservation struct {
	Site           string
	SampleDate     time.Time
	AnalysisDate   time.Time
	Analyte        string
	CasNum         string
	Result         float32
	Qualifier      string
	DetectionLimit float32
	Units          string
}

func ImportData() error {
	return errors.New("Not implemented")
}

func GetObservations() ([]PrecursorObservation, error) {
	return nil, errors.New("Not implemented")
}
