package mysql

import (
	"cmdTest/internal/dto/model"
)

type WellDAOMysql struct{}

func (w *WellDAOMysql) NewWell(brief *model.WellBrief, well *model.Well) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err := tx.Commit()
			if err != nil {
				_ = tx.Rollback()
				panic(err.Error())
			}
		}
	}()

	sqlStr := "INSERT INTO well_brief(well_name, well_type, note, create_time, update_time) VALUE (?, ?, ?, ?, ?)"
	exec, err := tx.Exec(
		sqlStr,
		brief.WellName,
		brief.WellType,
		brief.Note,
		brief.CreateTime,
		brief.UpdateTime,
	)
	if err != nil {
		return err
	}

	id, _ := exec.LastInsertId()

	sqlStr = "INSERT INTO well(ID, WELL_NAME, LIFE, ADDRESS, AFFILIATION, DEPTH, CONSTRUCTION, BOREHOLE_SIZE, FINISH_TIME, AVERAGE_DAILY_PRODUCTION) VALUE (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = tx.Exec(
		sqlStr,
		id,
		well.WellName,
		well.Life,
		well.Address,
		well.Affiliation,
		well.Depth,
		well.Construction,
		well.BoreholeSize,
		well.FinishTime,
		well.AverageDailyProduction,
	)
	return err
}

func (w *WellDAOMysql) ReviseWell(brief *model.WellBrief, well *model.Well) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err := tx.Commit()
			if err != nil {
				_ = tx.Rollback()
				panic(err.Error())
			}
		}
	}()

	sqlStr := "UPDATE well_brief SET well_name = ?, well_type = ?, note = ?, update_time = ? WHERE id = ?"
	_, err = tx.Exec(
		sqlStr,
		brief.WellName,
		brief.WellType,
		brief.Note,
		brief.UpdateTime,
		brief.ID,
	)
	if err != nil {
		return err
	}

	sqlStr = "UPDATE well SET WELL_NAME = ?, LIFE = ?, ADDRESS = ?, AFFILIATION = ?, DEPTH = ?, CONSTRUCTION = ?, BOREHOLE_SIZE = ?, FINISH_TIME = ?, AVERAGE_DAILY_PRODUCTION = ? WHERE id = ?"
	_, err = tx.Exec(
		sqlStr,
		well.WellName,
		well.Life,
		well.Address,
		well.Affiliation,
		well.Depth,
		well.Construction,
		well.BoreholeSize,
		well.FinishTime,
		well.AverageDailyProduction,
		well.ID,
	)
	return err
}

func (w *WellDAOMysql) DeleteWell(wellID string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err := tx.Commit()
			if err != nil {
				_ = tx.Rollback()
				panic(err.Error())
			}
		}
	}()

	sqlStr := "DELETE FROM well_brief WHERE id = ?"
	_, err = tx.Exec(sqlStr, wellID)
	if err != nil {
		return err
	}

	sqlStr = "DELETE FROM well WHERE id = ?"
	_, err = tx.Exec(sqlStr, wellID)
	return err
}

func (w *WellDAOMysql) GetBriefWellInfo() (infos []model.WellBrief, err error) {
	sqlStr := "SELECT * FROM well_brief"
	err = db.Select(&infos, sqlStr)
	return
}

func (w *WellDAOMysql) GetWellInfo(wellID string) (well *model.Well, err error) {
	well = &model.Well{}
	sqlStr := "SELECT * FROM well WHERE id = ?"
	err = db.Get(well, sqlStr, wellID)
	return
}
