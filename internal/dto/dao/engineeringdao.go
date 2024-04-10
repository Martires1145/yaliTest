package dao

import "cmdTest/internal/dto/model"

type EngineeringDao interface {
	NewEngineering(brief *model.EngineeringBrief, engineering *model.Engineering) (int64, error)
	NewDevice(devices []model.Device) error
	ReviseEngineering(brief *model.EngineeringBrief, engineering *model.Engineering) error
	AddDevices(devices []model.Device) error
	DeleteDevices(IDs []string) error
	DeleteEngineering(id string) error
	GetBriefEngineeringInfos() ([]model.EngineeringBrief, error)
	GetEngineeringInfo(id string) (*model.Engineering, error)
}
