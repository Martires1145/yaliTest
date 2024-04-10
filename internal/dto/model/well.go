package model

type WellBrief struct {
	ID         uint   `json:"ID" db:"id"`
	WellName   string `json:"wellName" db:"well_name"`
	WellType   string `json:"wellType" db:"well_type"`
	Note       string `json:"note" db:"note"`
	CreateTime int64  `json:"createTime" db:"create_time"`
	UpdateTime int64  `json:"updateTime" db:"update_time"`
}

type Well struct {
	ID                     uint   `json:"ID" db:"id"`
	WellName               string `json:"wellName" db:"well_name"`
	Life                   int    `json:"life" db:"life"`
	Address                string `json:"address" db:"address"`
	Affiliation            string `json:"affiliation" db:"affiliation"`
	Depth                  int    `json:"depth" db:"depth"`
	Construction           int    `json:"construction" db:"construction"`
	BoreholeSize           int    `json:"boreholeSize" db:"borehole_size"`
	FinishTime             int64  `json:"finishTime" db:"finish_time"`
	AverageDailyProduction int    `json:"averageDailyProduction" db:"average_daily_production"`
}

type WellJson struct {
	ID                     uint   `json:"ID" db:"id"`
	WellName               string `json:"wellName" db:"well_name"`
	WellType               string `json:"wellType" db:"well_type"`
	Note                   string `json:"note" db:"note"`
	Life                   int    `json:"life" db:"life"`
	Address                string `json:"address" db:"address"`
	Affiliation            string `json:"affiliation" db:"affiliation"`
	Depth                  int    `json:"depth" db:"depth"`
	Construction           int    `json:"construction" db:"construction"`
	BoreholeSize           int    `json:"boreholeSize" db:"borehole_size"`
	FinishTime             int64  `json:"finishTime" db:"finish_time"`
	AverageDailyProduction int    `json:"averageDailyProduction" db:"average_daily_production"`
}

func (w *WellJson) ToDBData() (WellBrief, Well) {
	return WellBrief{
			ID:       w.ID,
			WellName: w.WellName,
			WellType: w.WellType,
			Note:     w.Note,
		}, Well{
			ID:                     w.ID,
			WellName:               w.WellName,
			Life:                   w.Life,
			Address:                w.Address,
			Affiliation:            w.Affiliation,
			Depth:                  w.Depth,
			Construction:           w.Construction,
			BoreholeSize:           w.BoreholeSize,
			FinishTime:             w.FinishTime,
			AverageDailyProduction: w.AverageDailyProduction,
		}
}
