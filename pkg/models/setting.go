package models

type (
	Setting struct {
		ID        uint64 `json:"id" gorm:"primary_key"`
		Parameter string `json:"parameter"`
		Value     string `json:"value"`
	}
)

func (Setting) TableName() string {
	return "setting"
}

type Settings struct {
	Dev_ID   string `json:"device_id"`
	Interval string `json:"interval"`
	Broker   string `json:"broker"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
}
