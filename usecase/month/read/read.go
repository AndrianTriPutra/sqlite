package read

import (
	"context"
	"fmt"
	"log"
	"time"
	"xsqlitex/pkg/models"
	"xsqlitex/pkg/repository/month"
	"xsqlitex/pkg/repository/setting"
)

func Read(ctx context.Context, monthRepo month.IMonthRepository, settingRepo setting.ISettingRepository) {
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
	ymd := fmt.Sprintf("%02d-%02d-%02d", ts.Year(), ts.Month(), ts.Day())
	ymd += "%"
	ymdh := fmt.Sprintf("%02d-%02d-%02dT%02d", ts.Year(), ts.Month(), ts.Day(), ts.Hour())
	ymdh += "%"
	log.Printf("device_id:%s", Setting.Dev_ID)
	log.Printf("sensorid :%v", Setting.Sensor_ID)
	log.Printf("month    :%s", month)
	log.Printf("ymd      :%s", ymd)
	log.Printf("ymdh     :%s", ymdh)

	var result []models.Month
	data, err := monthRepo.ReadDataBy(ctx, Setting.Dev_ID, Setting.Sensor_ID, month, ymdh) // 1 hour
	//data, err := monthRepo.ReadDataBy(ctx, Setting.Dev_ID, Setting.Sensor_ID, month, ymd) // 1 day
	if err != nil {
		log.Fatalf("error read data:%v", err)
	}

	for _, val := range data {
		result = append(result, val)
	}

	i := 1
	for _, val := range result {
		fmt.Printf("[%v] ", i)
		fmt.Println(val)
		i++
	}
}
