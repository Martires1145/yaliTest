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

	err = modelDaoMysql.NewModel(modelData, id)
	return err
}

func DeleteModel(id string) error {
	err := modelDaoMysql.DeleteModel(id)
	return err
}
