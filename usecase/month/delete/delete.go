package delete

import (
	"context"
	"fmt"
	"log"
	"time"
	"xsqlitex/pkg/models"
	"xsqlitex/pkg/repository/month"
	"xsqlitex/pkg/repository/setting"
)

func Deleted(ctx context.Context, monthRepo month.IMonthRepository, settingRepo setting.ISettingRepository) {
	var Setting models.Month

	setting, err := settingRepo.ReadAll(ctx)
	if err != nil {
		log.Fatalf("error read setting : %s", err)
	}

	for _, val := range setting {
		if val.Parameter == "device_id" {
			Setting.Dev_ID = val.Value
		}
	}

	Setting.Sensor_ID = uint64(1)
	ts := time.Now()
	month := fmt.Sprintf("%02d", ts.Month())
	log.Printf("device_id:%s", Setting.Dev_ID)
	log.Printf("sensorid :%v", Setting.Sensor_ID)
	log.Printf("month    :%s", month)

	err = monthRepo.DeleteDataByMonth(ctx, Setting.Dev_ID, Setting.Sensor_ID, month)
	if err != nil {
		log.Fatalf("error delete data:%v", err)
	}
	log.Println("succes delete")
}
