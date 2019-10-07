package model

import (
	precursorModel "co-air-quality.api/src/precursor/model"
)

func ImportData() error {
	var err = precursorModel.ImportData()
	return err
}
