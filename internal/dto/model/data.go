package model

import (
	"encoding/json"
	"strconv"
)

type Data struct {
	Date          string `json:"date" csv:"date"`
	YaliStage     string `json:"yali_stage" csv:"yali_stage"`
	Stage         string `json:"stage" csv:"stage"`
	Taoya         string `json:"taoya" csv:"taoya"`
	Paichu        string `json:"paichu" csv:"paichu"`
	JieduanPaichu string `json:"jieduan_paichu" csv:"jieduan_paichu"`
	Paichu1       string `json:"paichu_1" csv:"paichu_1"`
	Shanongdu     string `json:"shanongdu" csv:"shanongdu"`
	Press         string `json:"press" csv:"press"`
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
	Max  float64   `json:"maxYali"`
	Min  float64   `json:"min"`
	Mean float64   `json:"mean"`
	Yali []float64 `json:"yali"`
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

func (d *Data) ParseTime() int64 {
	// todo
	t, _ := strconv.Atoi(d.Date)
	return int64(t)

}
