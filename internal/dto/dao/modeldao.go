package dao

import "cmdTest/internal/dto/model"

type ModelDao interface {
	NewModel(modelData *model.JsonModel, ParamsID int64) error
	NewParams(paramsJson *model.ParamsJson) (ID int64, err error)
	DeleteModel(id string) error
}
