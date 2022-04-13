package insert

import (
	"context"
	"log"
	"xsqlitex/pkg/models"
	"xsqlitex/pkg/repository/setting"
)

func Insert(ctx context.Context, settingRepo setting.ISettingRepository) {
	mde_map(ctx, settingRepo)
	mde_arr(ctx, settingRepo)
}

func mde_map(ctx context.Context, settingRepo setting.ISettingRepository) {
	var config models.Setting

	var settings = map[string]string{
		"device_id": "dev_001",
		"interval":  "1",
		"broker":    "broker.emqx.io:1883",
		"user":      "emqx",
		"pass":      "public",
	}

	for key, val := range settings {
		log.Printf("%s : %s", key, val)
		config.Parameter = key
		config.Value = val

		err := settingRepo.InsertDatatoTable(ctx, config)
		if err != nil {
			log.Fatalf("error insert table setting :%s", err)
		}
		log.Println("succes insert table setting")

	}
}

func mde_arr(ctx context.Context, settingRepo setting.ISettingRepository) {
	var config models.Setting

	var paramater [5]string
	paramater[0] = "device_id"
	paramater[1] = "interval"
	paramater[2] = "broker"
	paramater[3] = "user"
	paramater[4] = "pass"

	var value [5]string
	value[0] = "dev_001"
	value[1] = "1"
	value[2] = "broker.emqx.io:1883"
	value[3] = "emqx"
	value[4] = "public"

	for i := 0; i < 5; i++ {
		config.Parameter = paramater[i]
		config.Value = value[i]
		log.Println(config)
		err := settingRepo.InsertDatatoTable(ctx, config)
		if err != nil {
			log.Fatalf("error insert table setting :%s", err)
		}
		log.Println("succes insert table setting")
	}
}
