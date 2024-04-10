package server

import (
	"cmdTest/internal/dto/model"
	"cmdTest/internal/dto/mysql"
	"time"
)

var engineeringDao = mysql.EngineeringDaoMysql{}

func NewEngineering(engineeringData *model.EngineeringJson) error {
	brief, engineering := engineeringData.ToDBData()
	brief.CreateTime, brief.UpdateTime = time.Now().Unix(), time.Now().Unix()

	id, err := engineeringDao.NewEngineering(&brief, &engineering)
	if err != nil {
		return err
	}

	for i := range engineering.Devices {
		engineering.Devices[i].EngineeringID = uint(id)
	}

	return engineeringDao.NewDevice(engineering.Devices)
}

func ReviseEngineering(engineeringData *model.EngineeringJson) error {
	brief, engineering := engineeringData.ToDBData()
	brief.UpdateTime = time.Now().Unix()

	return engineeringDao.ReviseEngineering(&brief, &engineering)
}

func AddDevices(devices []model.Device) error {
	return engineeringDao.AddDevices(devices)
}

func DeleteDevices(IDs []string) error {
	return engineeringDao.DeleteDevices(IDs)
}

func DeleteEngineering(id string) error {
	return engineeringDao.DeleteEngineering(id)
}

func GetBriefEngineeringInfos() ([]model.EngineeringBrief, error) {
	return engineeringDao.GetBriefEngineeringInfos()
}

func GetEngineeringInfo(id string) (*model.Engineering, error) {
	return engineeringDao.GetEngineeringInfo(id)
}
