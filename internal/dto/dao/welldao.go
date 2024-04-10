package dao

import "cmdTest/internal/dto/model"

type WellDao interface {
	NewWell(brief *model.WellBrief, well *model.Well) error
	ReviseWell(brief *model.WellBrief, well *model.Well) error
	DeleteWell(wellID string) error
	GetBriefWellInfo() ([]model.WellBrief, error)
	GetWellInfo(wellID string) (*model.Well, error)
}
