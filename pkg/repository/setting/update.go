package setting

import (
	"context"
	"xsqlitex/pkg/models"

	"gorm.io/gorm"
)

func (r *settingRepository) Update(ctx context.Context, key string, data models.Setting) error {
	db := r.provider.Db(ctx).(*gorm.DB)
	var err error

	err = db.Table(models.Setting{}.TableName()).Where("parameter = ?", key).Updates(&data).Error
	return err
}
