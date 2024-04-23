package dao

import "cmdTest/internal/dto/model"

type ModelDao interface {
	NewModel(modelData *model.JsonModel, ParamsID int) error
	NewParams(paramsJson *model.ParamsJson) (ID int, err error)
}
