package setting

import (
	"context"
	"xsqlitex/pkg/models"

	"gorm.io/gorm"
)

func (r *settingRepository) ReadAll(ctx context.Context) ([]models.Setting, error) {
	db := r.provider.Db(ctx).(*gorm.DB)
	var data []models.Setting
	var err error
	err = db.Table(models.Setting{}.TableName()).Find(&data).Error

	return data, err
}
