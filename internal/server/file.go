package server

import (
	"cmdTest/internal/dto/model"
	"cmdTest/internal/dto/mysql"
	"cmdTest/pkg/util"
	"github.com/spf13/viper"
	"mime/multipart"
	"os"
)

var (
	basePath     = viper.GetString("model.path")
	fileDaoMysql = mysql.FileDaoMySql{}
)

func NewFilePath(path string) error {
	newPath := basePath + "\\" + path

	err := os.Mkdir(newPath, os.ModePerm)
	if err != nil {
		return err
	}

	return fileDaoMysql.NewFile(path)
}

func GetAllPath() ([]model.FileData, error) {
	return fileDaoMysql.GetAllPath()
}

func GetFile(path string) ([]string, error) {
	path = basePath + "\\" + path
	return util.GetFile(path)
}

func SaveCsvFile(path string, fileData *multipart.FileHeader) error {
	path = basePath + "\\" + path + "\\" + fileData.Filename
	return util.SaveFile(path, fileData)
}

func DeletePath(path string) error {
	err := fileDaoMysql.DeletePath(path)
	if err != nil {
		return err
	}

	path = basePath + "\\" + path
	return util.DeleteFile(path)
}

func DeleteFile(path, name string) error {
	path = basePath + "\\" + path + "\\" + name
	return util.DeleteFile(path)
}
