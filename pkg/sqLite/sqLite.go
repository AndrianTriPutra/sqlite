package sqLite

import (
	"context"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"xsqlitex/pkg/models"
)

type sqliteDB struct {
	dsn string
	db  *gorm.DB
}

func NewsqliteDb(dsn string) IDatabase {
	return &sqliteDB{
		dsn: dsn,
	}
}

func (p *sqliteDB) Connect(ctx context.Context) error {
	var err error
	p.db, err = gorm.Open(sqlite.Open(p.dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	p.db.AutoMigrate(&models.Month1{}, &models.Month2{}, &models.Month3{}, &models.Month4{},
		&models.Month5{}, &models.Month6{}, &models.Month7{}, &models.Month8{},
		&models.Month9{}, &models.Month10{}, &models.Month11{}, &models.Month12{},
		&models.Setting{})
	log.Printf("successfully connect to database: %v", p.dsn)
	return nil
}

func (p sqliteDB) Db(ctx context.Context) interface{} {
	tx := ctx.Value("txContext")
	if tx == nil {
		return p.db
	}
	return tx.(*gorm.DB)
}
