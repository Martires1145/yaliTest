package dao

import "cmdTest/internal/dto/model"

type DataDao interface {
	SaveForecastDataHistory(id string, trueDataPath string) error
	GetHistoryData() ([]model.DataHistoryJson, error)
	DeleteHistoryData(id string) (string, string, error)
	GetOneHistory(id string) (*model.DataHistory, error)
}
