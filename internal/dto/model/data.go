package model

import "encoding/json"

type Data struct {
	Date          string `json:"date"`
	YaliStage     string `json:"yali_stage"`
	Stage         string `json:"stage"`
	Taoya         string `json:"taoya"`
	Paichu        string `json:"paichu"`
	JieduanPaichu string `json:"jieduan_paichu"`
	Paichu1       string `json:"paichu_1"`
	Shanongdu     string `json:"shanongdu"`
	Press         string `json:"press"`
}

type DataHistory struct {
	ID            uint   `json:"ID" db:"id"`
	ModelID       uint   `json:"modelID" db:"model_id"`
	WellID        uint   `json:"wellID" db:"well_id"`
	EngineeringID uint   `json:"engineeringID" db:"engineering_id"`
	CreateTime    int64  `json:"createTime" db:"create_time"`
	TrueDataPath  string `json:"trueDataPath" db:"true_data_path"`
	PDataPath     string `json:"pDataPath" db:"p_data_path"`
}

type DataHistoryJson struct {
	ID            uint  `json:"ID" db:"id"`
	ModelID       uint  `json:"modelID" db:"model_id"`
	WellID        uint  `json:"wellID" db:"well_id"`
	EngineeringID uint  `json:"engineeringID" db:"engineering_id"`
	CreateTime    int64 `json:"createTime" db:"create_time"`
}

type RangeData struct {
	Max  int    `json:"maxYali"`
	Min  int    `json:"min"`
	Mean int    `json:"mean"`
	Data []Data `json:"data"`
}

func (d *Data) Json() ([]byte, error) {
	bytes, err := json.Marshal(d)
	return bytes, err
}

func MakeData(list []string) *Data {
	return &Data{
		Date:          list[0],
		YaliStage:     list[1],
		Stage:         list[2],
		Taoya:         list[3],
		Paichu:        list[4],
		JieduanPaichu: list[5],
		Paichu1:       list[6],
		Shanongdu:     list[7],
		Press:         list[8],
	}
}
