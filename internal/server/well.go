package server

import (
	"cmdTest/internal/dto/model"
	"cmdTest/internal/dto/mysql"
	"time"
)

var wellDao = mysql.WellDAOMysql{}

func NewWell(wellData *model.WellJson) error {
	brief, well := wellData.ToDBData()
	brief.CreateTime, brief.UpdateTime = time.Now().Unix(), time.Now().Unix()

	return wellDao.NewWell(&brief, &well)
}

func ReviseWell(wellData *model.WellJson) error {
	brief, well := wellData.ToDBData()
	brief.UpdateTime = time.Now().Unix()

	return wellDao.ReviseWell(&brief, &well)
}

func DeleteWell(id string) error {
	return wellDao.DeleteWell(id)
}

func GetBriefWellInfo() ([]model.WellBrief, error) {
	return wellDao.GetBriefWellInfo()
}

func GetWellInfo(id string) (*model.Well, error) {
	return wellDao.GetWellInfo(id)
}
