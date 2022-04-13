package update

import (
	"context"
	"log"
	"xsqlitex/pkg/models"
	"xsqlitex/pkg/repository/setting"
)

func Update(ctx context.Context, settingRepo setting.ISettingRepository) {

	var data models.Setting
	data.Parameter = "pass"
	data.Value = "private"

	err := settingRepo.Update(ctx, data.Parameter, data)
	if err != nil {
		log.Fatalf("error update setting :%s", err)
	}
	log.Println("succ update setting")

}
