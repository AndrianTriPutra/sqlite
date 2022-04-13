package month

import (
	"context"
	"xsqlitex/pkg/models"
	"xsqlitex/pkg/sqLite"
)

type IMonthRepository interface {
	InsertDatatoTable(ctx context.Context, month string, data models.Month) error
	ReadDataBy(ctx context.Context, loc string, sensorid uint64, month, timestamp string) ([]models.Month, error)
	DeleteDataByMonth(ctx context.Context, loc string, sensorid uint64, month string) error
}

type monthRepository struct {
	provider sqLite.IDatabase
}

func NewRepository(db sqLite.IDatabase) IMonthRepository {
	return &monthRepository{
		provider: db,
	}
}
