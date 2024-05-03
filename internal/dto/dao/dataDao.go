package dao

import "cmdTest/internal/dto/model"

type DataDao interface {
	SaveDataHistory(history *model.DataHistory) error
	GetHistoryData() ([]model.DataHistoryJson, error)
	DeleteHistoryData(id string) (string, string, error)
	GetOneHistory(id string) (*model.DataHistory, error)
}
