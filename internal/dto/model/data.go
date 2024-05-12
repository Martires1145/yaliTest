package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type Data struct {
	Date          string `json:"date" csv:"date"`
	YaliStage     string `json:"yali_stage" csv:"yaliestage"`
	Stage         string `json:"stage" csv:"stage"`
	Taoya         string `json:"taoya" csv:"taoya"`
	Paichu        string `json:"paichu" csv:"paichu"`
	JieduanPaichu string `json:"jieduan_paichu" csv:"jieduanpaich"`
	Paichu1       string `json:"paichu_1" csv:"paichu.1"`
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
	Max      float64 `json:"maxYali"`
	Min      float64 `json:"min"`
	Mean     float64 `json:"mean"`
	Variance float64 `json:"variance"`
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

func (d *Data) ParseTime(day string) int64 {
	const layout = "2006-01-02 15:04:05"

	// 使用time.Parse解析时间字符串
	t, err := time.Parse(layout, day+d.Date)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return 0
	}
	return t.Unix()
}

func (d DataHistory) GetDay() string {
	t := time.Unix(d.CreateTime, 0)

	// 使用Format方法将时间格式化为"2006-01-02"
	formattedDate := t.Format("2006-01-02")

	return formattedDate + " "
}
