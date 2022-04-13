package month

import (
	"context"
	"xsqlitex/pkg/models"

	"gorm.io/gorm"
)

func (r *monthRepository) ReadDataBy(ctx context.Context, dev string, sensorid uint64, month, timestamp string) ([]models.Month, error) {
	db := r.provider.Db(ctx).(*gorm.DB)
	var data []models.Month
	var err error

	switch month {
	case "01":
		err = db.Table(models.Month1{}.TableName()).Where("dev_id = ? AND sensor_id = ? AND timestamp LIKE ?", dev, sensorid, timestamp).Find(&data).Error
	case "02":
		err = db.Table(models.Month2{}.TableName()).Where("dev_id = ? AND sensor_id = ? AND timestamp LIKE ?", dev, sensorid, timestamp).Find(&data).Error
	case "03":
		err = db.Table(models.Month3{}.TableName()).Where("dev_id = ? AND sensor_id = ? AND timestamp LIKE ?", dev, sensorid, timestamp).Find(&data).Error
	case "04":
		err = db.Table(models.Month4{}.TableName()).Where("dev_id = ? AND sensor_id = ? AND timestamp LIKE ?", dev, sensorid, timestamp).Find(&data).Error
	case "05":
		err = db.Table(models.Month5{}.TableName()).Where("dev_id = ? AND sensor_id = ? AND timestamp LIKE ?", dev, sensorid, timestamp).Find(&data).Error
	case "06":
		err = db.Table(models.Month6{}.TableName()).Where("dev_id = ? AND sensor_id = ? AND timestamp LIKE ?", dev, sensorid, timestamp).Find(&data).Error
	case "07":
		err = db.Table(models.Month7{}.TableName()).Where("dev_id = ? AND sensor_id = ? AND timestamp LIKE ?", dev, sensorid, timestamp).Find(&data).Error
	case "08":
		err = db.Table(models.Month8{}.TableName()).Where("dev_id = ? AND sensor_id = ? AND timestamp LIKE ?", dev, sensorid, timestamp).Find(&data).Error
	case "09":
		err = db.Table(models.Month9{}.TableName()).Where("dev_id = ? AND sensor_id = ? AND timestamp LIKE ?", dev, sensorid, timestamp).Find(&data).Error
	case "10":
		err = db.Table(models.Month10{}.TableName()).Where("dev_id = ? AND sensor_id = ? AND timestamp LIKE ?", dev, sensorid, timestamp).Find(&data).Error
	case "11":
		err = db.Table(models.Month11{}.TableName()).Where("dev_id = ? AND sensor_id = ? AND timestamp LIKE ?", dev, sensorid, timestamp).Find(&data).Error
	case "12":
		err = db.Table(models.Month12{}.TableName()).Where("dev_id = ? AND sensor_id = ? AND timestamp LIKE ?", dev, sensorid, timestamp).Find(&data).Error
	}

	return data, err
}
