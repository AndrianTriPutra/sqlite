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
