package convert

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"xsqlitex/pkg/models"
	"xsqlitex/pkg/repository/month"
	"xsqlitex/pkg/repository/setting"
)

func ExportoCSV(ctx context.Context, monthRepo month.IMonthRepository, settingRepo setting.ISettingRepository) {
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
	Setting.Sensor_ID = uint64(2)
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

	data, err := monthRepo.ReadDataBy(ctx, Setting.Dev_ID, Setting.Sensor_ID, month, ymd)
	if err != nil {
		log.Fatalf("error read data:%v", err)
	}

	filename := Setting.Dev_ID + ".csv"
	csvFile, err := os.Create(filename)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)

	no := 0
	var value [][]string
	var header = []string{"no", "id", "dev", "sensorid", "timestamp", "temp", "rh"}
	err = csvwriter.Write(header)
	if err != nil {
		log.Fatalf("failed write header: %s", err)
	}
	for _, record := range data {

		no++
		numb := strconv.Itoa(no)
		id := strconv.Itoa(int(record.ID))
		sensorid := strconv.Itoa(int(record.Sensor_ID))
		temp := fmt.Sprintf("%.2f", record.Temp)
		rh := fmt.Sprintf("%.2f", record.Rh)

		row := []string{numb, id, record.Dev_ID, sensorid, record.Timestamp, temp, rh}
		value = append(value, row)

	}

	err = csvwriter.WriteAll(value)
	if err != nil {
		log.Fatalf("failed write body: %s", err)
	}
	log.Println("succes write csv")
}
