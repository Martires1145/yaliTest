package mysql

import (
	"cmdTest/internal/dto/model"
	"github.com/jmoiron/sqlx"
)

type EngineeringDaoMysql struct{}

func (e *EngineeringDaoMysql) NewEngineering(brief *model.EngineeringBrief, engineering *model.Engineering) (id int64, err error) {
	tx, err := db.Begin()
	if err != nil {
		return -1, err
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

	sqlStr := "INSERT INTO engineering_brief(name, progress, state, create_time, update_time) VALUE (?, ?, ?, ?, ?)"
	exec, err := tx.Exec(
		sqlStr,
		brief.Name,
		brief.Progress,
		brief.State,
		brief.CreateTime,
		brief.UpdateTime,
	)
	if err != nil {
		return -1, err
	}
	id, _ = exec.LastInsertId()

	sqlStr = "INSERT INTO engineering(id, name, construction_unit, wellName, address, head, state, number_of_constructors, begin_time, estimated_completion_time) VALUE (?, ?, ?, ?, ?,?, ?, ?, ?, ?)"
	_, err = tx.Exec(
		sqlStr,
		id,
		engineering.Name,
		engineering.ConstructionUnit,
		engineering.WellName,
		engineering.Address,
		engineering.Head,
		engineering.State,
		engineering.NumberOfConstructors,
		engineering.BeginTime,
		engineering.EstimatedCompletionTime,
	)
	if err != nil {
		return -1, err
	}
	return -1, err
}

func (e *EngineeringDaoMysql) NewDevice(devices []model.Device) error {
	sqlStr := "INSERT INTO devices(engineering_id, name_of_device, number_of_device, type_of_device) VALUES (:engineering_id, :name_of_device, :number_of_device, :type_of_device)"
	_, err := db.NamedExec(sqlStr, devices)
	return err
}

func (e *EngineeringDaoMysql) ReviseEngineering(brief *model.EngineeringBrief, engineering *model.Engineering) error {
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

	sqlStr := "UPDATE engineering_brief SET name = ?, progress = ?, state = ?, update_time = ? WHERE id = ?"
	_, err = tx.Exec(
		sqlStr,
		brief.Name,
		brief.Progress,
		brief.State,
		brief.UpdateTime,
		brief.ID,
	)
	if err != nil {
		return err
	}

	sqlStr = "UPDATE engineering SET name = ?, construction_unit = ?, wellName = ?, address = ?, head = ?, state = ?, number_of_constructors = ?, begin_time = ?, estimated_completion_time = ? WHERE id = ?"
	_, err = tx.Exec(
		sqlStr,
		engineering.Name,
		engineering.ConstructionUnit,
		engineering.WellName,
		engineering.Address,
		engineering.Head,
		engineering.State,
		engineering.NumberOfConstructors,
		engineering.BeginTime,
		engineering.EstimatedCompletionTime,
		engineering.ID,
	)
	return err
}

func (e *EngineeringDaoMysql) AddDevices(devices []model.Device) error {
	sqlStr := "INSERT INTO devices (engineering_id, name_of_device, number_of_device, type_of_device) VALUES (:engineering_id, :name_of_device, :number_of_device, :type_of_device)"
	_, err := db.NamedExec(sqlStr, devices)
	return err
}

func (e *EngineeringDaoMysql) DeleteDevices(IDs []string) error {
	sqlStr := "DELETE FROM devices WHERE id IN (?)"
	query, args, err := sqlx.In(sqlStr, IDs)
	if err != nil {
		return err
	}

	_, err = db.Exec(query, args...)
	return err
}

func (e *EngineeringDaoMysql) DeleteEngineering(id string) error {
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

	sqlStr := "DELETE FROM engineering WHERE id = ?"
	_, err = tx.Exec(sqlStr, id)
	if err != nil {
		return err
	}

	sqlStr = "DELETE FROM engineering_brief WHERE id = ?"
	_, err = tx.Exec(sqlStr, id)
	if err != nil {
		return err
	}

	sqlStr = "DELETE FROM devices WHERE engineering_id = ?"
	_, err = tx.Exec(sqlStr, id)
	return err
}

func (e *EngineeringDaoMysql) GetBriefEngineeringInfos() (briefInfos []model.EngineeringBrief, err error) {
	sqlStr := "SELECT * FROM engineering_brief"
	err = db.Select(&briefInfos, sqlStr)
	return
}

func (e *EngineeringDaoMysql) GetEngineeringInfo(id string) (info *model.Engineering, err error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
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

	sqlStr := "SELECT * FROM engineering WHERE id = ?"
	rows, err := tx.Query(sqlStr, id)
	if err != nil {
		return nil, err
	}

	err = rows.Scan(&info)
	if err != nil {
		return nil, err
	}

	sqlStr = "SELECT * FROM devices WHERE engineering_id = ?"
	devices, err := tx.Query(sqlStr, id)
	if err != nil {
		return nil, err
	}

	err = devices.Scan(&info.Devices)
	return
}
