package setting

import (
	"context"
	"xsqlitex/pkg/models"

	"gorm.io/gorm"
)

func (r *settingRepository) InsertDatatoTable(ctx context.Context, data models.Setting) error {
	var err error
	db := r.provider.Db(ctx).(*gorm.DB)
	err = db.Table(models.Setting{}.TableName()).Create(&data).Error
	return err
}
