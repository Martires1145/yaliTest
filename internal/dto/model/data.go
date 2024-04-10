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

func (d *Data) Json() ([]byte, error) {
	bytes, err := json.Marshal(d)
	return bytes, err
}

func MakeData(list []string) Data {
	return Data{
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
