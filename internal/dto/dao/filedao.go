package dao

import "cmdTest/internal/dto/model"

type FileDao interface {
	NewFile(fileName string) error
	GetAllPath() ([]model.FileData, error)
	DeletePath(path string) error
}
