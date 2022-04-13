package month

import (
	"context"
	"xsqlitex/pkg/models"

	"gorm.io/gorm"
)

func (r *monthRepository) InsertDatatoTable(ctx context.Context, month string, data models.Month) error {
	var err error
	db := r.provider.Db(ctx).(*gorm.DB)

	switch month {
	case "01":
		err = db.Table(models.Month1{}.TableName()).Create(&data).Error
	case "02":
		err = db.Table(models.Month2{}.TableName()).Create(&data).Error
	case "03":
		err = db.Table(models.Month3{}.TableName()).Create(&data).Error
	case "04":
		err = db.Table(models.Month4{}.TableName()).Create(&data).Error
	case "05":
		err = db.Table(models.Month5{}.TableName()).Create(&data).Error
	case "06":
		err = db.Table(models.Month6{}.TableName()).Create(&data).Error
	case "07":
		err = db.Table(models.Month7{}.TableName()).Create(&data).Error
	case "08":
		err = db.Table(models.Month8{}.TableName()).Create(&data).Error
	case "09":
		err = db.Table(models.Month9{}.TableName()).Create(&data).Error
	case "10":
		err = db.Table(models.Month10{}.TableName()).Create(&data).Error
	case "11":
		err = db.Table(models.Month11{}.TableName()).Create(&data).Error
	case "12":
		err = db.Table(models.Month12{}.TableName()).Create(&data).Error
	}

	return err
}
