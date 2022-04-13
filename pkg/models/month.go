package models

type (
	Month struct {
		ID        uint64  `json:"id" gorm:"primary_key"`
		Dev_ID    string  `json:"dev_id"`
		Sensor_ID uint64  `json:"sensor_id"`
		Timestamp string  `json:"timestamp"`
		Temp      float32 `json:"temp"`
		Rh        float32 `json:"rh"`
	}

	Month1  Month
	Month2  Month
	Month3  Month
	Month4  Month
	Month5  Month
	Month6  Month
	Month7  Month
	Month8  Month
	Month9  Month
	Month10 Month
	Month11 Month
	Month12 Month
)

func (Month1) TableName() string {
	return "month01"
}

func (Month2) TableName() string {
	return "month02"
}

func (Month3) TableName() string {
	return "month03"
}

func (Month4) TableName() string {
	return "month04"
}

func (Month5) TableName() string {
	return "month05"
}

func (Month6) TableName() string {
	return "month06"
}

func (Month7) TableName() string {
	return "month07"
}

func (Month8) TableName() string {
	return "month08"
}

func (Month9) TableName() string {
	return "month09"
}

func (Month10) TableName() string {
	return "month10"
}

func (Month11) TableName() string {
	return "month11"
}

func (Month12) TableName() string {
	return "month12"
}
