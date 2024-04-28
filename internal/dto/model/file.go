package model

type FileData struct {
	ID       uint   `json:"id" db:"id"`
	FileName string `json:"fileName" db:"file_name"`
}
