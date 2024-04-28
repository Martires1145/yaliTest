package mysql

import (
	"cmdTest/internal/dto/model"
)

type FileDaoMySql struct{}

func (f *FileDaoMySql) NewFile(fileName string) error {
	sqlStr := "INSERT INTO files(file_name) VALUE(?)"

	_, err := db.Exec(sqlStr, fileName)
	return err
}

func (f *FileDaoMySql) GetAllPath() (paths []model.FileData, err error) {
	sqlStr := "SELECT * FROM files"

	err = db.Select(&paths, sqlStr)
	return
}

func (f *FileDaoMySql) DeletePath(path string) error {
	sqlStr := "DELETE FROM files WHERE file_name = ?"

	_, err := db.Exec(sqlStr, path)
	return err
}
