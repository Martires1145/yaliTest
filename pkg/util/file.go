package util

import (
	"github.com/spf13/viper"
	"os"
)

var (
	rootPath  = viper.GetString("model.rootPath")
	dataPath  = viper.GetString("model.dataPath")
	trainPath = viper.GetString("model.trainPath")
	valiPath  = viper.GetString("model.valiPath")
	testPath  = viper.GetString("model.testPath")
)

func GetFile(fileType string) (f []string, err error) {
	path := getPath(fileType)

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		f = append(f, path+file.Name())
	}

	return
}

func getPath(fileType string) string {
	var path string
	switch fileType {
	case "root":
		path = rootPath
	case "data":
		path = dataPath
	case "train":
		path = trainPath
	case "vali":
		path = valiPath
	case "test":
		path = testPath
	}
	return path
}
