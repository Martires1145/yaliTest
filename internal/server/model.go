package server

import (
	"cmdTest/internal/dto/model"
	"cmdTest/internal/dto/mysql"
)

var modelDaoMysql = mysql.ModelDaoMysql{}

func NewModel(modelData *model.JsonModel) error {
	modelData.SetTime()

	id, err := modelDaoMysql.NewParams(modelData.Params)
	if err != nil {
		return err
	}

	return modelDaoMysql.NewModel(modelData, id)
}

func DeleteModel(id string) error {
	return modelDaoMysql.DeleteModel(id)
}

func ModifyModel(id, name, useKafka string) error {
	return modelDaoMysql.ModifyModel(id, name, useKafka)
}

func CopyModel(id, name string) error {
	return modelDaoMysql.CopyModel(id, name)
}

func GetAllModel() ([]model.DBModel, error) {
	return modelDaoMysql.GetAllModel()
}

func GetModelParams(id string) (*model.ParamsJson, error) {
	params, err := modelDaoMysql.GetModelParams(id)
	if err != nil {
		return nil, err
	}

	params.SetUseExtra()
	return params, nil
}
