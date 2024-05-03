package dao

import "cmdTest/internal/dto/model"

type ModelDao interface {
	NewModel(modelData *model.JsonModel) (int64, error)
	NewParams(paramsJson *model.ParamsJson, modelID int64) (ID int64, err error)
	DeleteModel(id string) error
	ModifyModel(id, name, useKafka string) error
	CopyModel(id, name string) error
	GetAllModel() ([]model.DBModel, error)
	GetModelParams(id string) (*model.ParamsJson, error)
	UpdateModelParamsID(modelID, paramsID int64) error
	UseModel(id string) error
}
