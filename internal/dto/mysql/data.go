package mysql

import (
	"cmdTest/internal/dto/model"
)

type DataDaoMysql struct{}

func (d *DataDaoMysql) SaveForecastDataHistory(id string, trueDataPath string) error {
	sqlStr := "UPDATE histories SET p_data_path = ? WHERE id = ?"
	_, err := db.Exec(sqlStr, trueDataPath, id)
	return err
}

func (d *DataDaoMysql) GetHistoryData() (histories []model.DataHistoryJson, err error) {
	sqlStr := "SELECT id, model_id, well_id, engineering_id, create_time FROM histories"
	err = db.Select(&histories, sqlStr)
	return
}

func (d *DataDaoMysql) DeleteHistoryData(id string) (t string, f string, err error) {
	var history model.DataHistory
	sqlStr := "SELECT  * FROM histories WHERE id = ?"
	err = db.Get(&history, sqlStr, id)
	if err != nil {
		return "", "", err
	}

	sqlStr = "DELETE FROM histories WHERE id = ?"
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		return "", "", err
	}

	return history.TrueDataPath, history.PDataPath, nil
}

func (d *DataDaoMysql) GetOneHistory(id string) (history *model.DataHistory, err error) {
	sqlStr := "SELECT * FROM histories WHERE id = ?"
	err = db.Get(history, sqlStr, id)
	return
}
