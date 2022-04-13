package repository

import (
	"xsqlitex/pkg/repository/month"
	"xsqlitex/pkg/repository/setting"
	"xsqlitex/pkg/sqLite"
)

func Init(db sqLite.IDatabase) (month.IMonthRepository, setting.ISettingRepository) {
	monthRepo := month.NewRepository(db)
	settingRepo := setting.NewRepository(db)

	return monthRepo, settingRepo
}
