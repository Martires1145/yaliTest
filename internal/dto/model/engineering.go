package model

type EngineeringBrief struct {
	ID         uint    `json:"ID" db:"id"`
	Name       string  `json:"name" db:"name"`
	Progress   float64 `json:"progress" db:"progress"`
	State      int     `json:"state" db:"state"`
	CreateTime int64   `json:"createTime" db:"create_time"`
	UpdateTime int64   `json:"updateTime" db:"update_time"`
}

type Engineering struct {
	ID                      uint     `json:"ID" db:"id"`
	Name                    string   `json:"name" db:"name"`
	ConstructionUnit        string   `json:"constructionUnit" db:"construction_unit"`
	WellName                string   `json:"wellName" db:"wellName"`
	Address                 string   `json:"address" db:"address"`
	Head                    string   `json:"head" db:"head"`
	State                   int      `json:"state" db:"state"`
	NumberOfConstructors    int      `json:"numberOfConstructors" db:"number_of_constructors"`
	BeginTime               int64    `json:"beginTime" db:"begin_time"`
	EstimatedCompletionTime int64    `json:"estimatedCompletionTime" db:"estimated_completion_time"`
	Devices                 []Device `json:"devices"`
}

type Device struct {
	ID             uint   `json:"ID" db:"id"`
	EngineeringID  uint   `json:"engineeringID" db:"engineering_id"`
	NameOfDevice   string `json:"nameOfDevice" db:"name_of_device"`
	NumberOfDevice int    `json:"numberOfDevice" db:"number_of_device"`
	TypeOfDevice   string `json:"typeOfDevice" db:"type_of_device"`
}

type EngineeringJson struct {
	ID                      uint     `json:"ID" db:"id"`
	Name                    string   `json:"name" db:"name"`
	Progress                float64  `json:"progress" db:"progress"`
	State                   int      `json:"state" db:"state"`
	ConstructionUnit        string   `json:"constructionUnit" db:"construction_unit"`
	WellName                string   `json:"wellName" db:"wellName"`
	Address                 string   `json:"address" db:"address"`
	Head                    string   `json:"head" db:"head"`
	NumberOfConstructors    int      `json:"numberOfConstructors" db:"number_of_constructors"`
	BeginTime               int64    `json:"beginTime" db:"begin_time"`
	EstimatedCompletionTime int64    `json:"estimatedCompletionTime" db:"estimated_completion_time"`
	Devices                 []Device `json:"devices"`
}

func (e *EngineeringJson) ToDBData() (EngineeringBrief, Engineering) {
	return EngineeringBrief{
			ID:       e.ID,
			Name:     e.Name,
			Progress: e.Progress,
			State:    e.State,
		}, Engineering{
			ID:                      e.ID,
			Name:                    e.Name,
			ConstructionUnit:        e.ConstructionUnit,
			WellName:                e.WellName,
			Address:                 e.Address,
			Head:                    e.Head,
			State:                   e.State,
			NumberOfConstructors:    e.NumberOfConstructors,
			BeginTime:               e.BeginTime,
			EstimatedCompletionTime: e.EstimatedCompletionTime,
			Devices:                 e.Devices,
		}
}
