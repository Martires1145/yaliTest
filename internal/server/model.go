package server

import (
	"cmdTest/internal/dto/model"
	"cmdTest/internal/dto/mysql"
	"cmdTest/pkg/util"
	"github.com/spf13/viper"
	"mime/multipart"
)

var modelDaoMysql = mysql.ModelDaoMysql{}
var modelPath = viper.GetString("model.path")
var checkpointsPath = viper.GetString("model.checkpoints")

func NewModel(modelData *model.JsonModel) error {
	modelData.SetTime()

	modelID, err := modelDaoMysql.NewModel(modelData)

	id, err := modelDaoMysql.NewParams(modelData.Params, modelID)
	if err != nil {
		return err
	}

	return modelDaoMysql.UpdateModelParamsID(modelID, id)
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

func SavePthFile(id string, file *multipart.FileHeader) error {
	params, err := GetModelParams(id)
	if err != nil {
		return err
	}

	path := modelPath + checkpointsPath + util.MakeModelPath(params)

	return util.SaveFile(path, file)
}

func UseModel(id string) (isStream bool, clientID string, err error) {
	params, err := modelDaoMysql.GetModelParams(id)
	if err != nil {
		return false, "", err
	}

	isStream, err = Call(params.GetParams())
	if err != nil {
		return false, "", err
	}

	clientID = util.GetClientID()
	Conns[clientID] = nil

	err = modelDaoMysql.UseModel(id)
	if err != nil {
		return false, "", err
	}
	return
}
