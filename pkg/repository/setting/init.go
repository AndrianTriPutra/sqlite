package setting

import (
	"context"
	"xsqlitex/pkg/models"
	"xsqlitex/pkg/sqLite"
)

type ISettingRepository interface {
	InsertDatatoTable(ctx context.Context, data models.Setting) error
	ReadAll(ctx context.Context) ([]models.Setting, error)
	Update(ctx context.Context, key string, data models.Setting) error
}

type settingRepository struct {
	provider sqLite.IDatabase
}

func NewRepository(db sqLite.IDatabase) ISettingRepository {
	return &settingRepository{
		provider: db,
	}
}
