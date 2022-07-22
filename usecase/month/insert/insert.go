package insert

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"sync"
	"syscall"
	"time"

	"xsqlitex/pkg/models"
	"xsqlitex/pkg/repository/month"
	"xsqlitex/pkg/repository/setting"
)

func Insert(ctx context.Context, cancel context.CancelFunc, monthRepo month.IMonthRepository, settingRepo setting.ISettingRepository) {
	var wg sync.WaitGroup
	stoped := make(chan os.Signal, 1)
	signal.Notify(stoped,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	//log.Println(" insert 2 value to 1 table")
	var settingan models.Settings

	data, err := settingRepo.ReadAll(ctx)
	if err != nil {
		log.Fatalf("error read setting : %s", err)
	}

	for _, val := range data {
		switch val.Parameter {
		case "device_id":
			settingan.Dev_ID = val.Value
		case "interval":
			settingan.Interval = val.Value
		}
	}
	intInterval, _ := strconv.Atoi(settingan.Interval)
	interval := time.Duration(intInterval)
	interval = interval * time.Second
	log.Printf("device_id:%s", settingan.Dev_ID)
	log.Printf("interval :%v", interval)

	wg.Add(3)
	go func() {
		wg.Wait()
		cancel()
	}()

	go func() {
		defer wg.Done()

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		err := sensor1_TableMonth(*ticker, ctx, monthRepo)
		if err != nil {
			log.Fatalf("Error sensor1_TableMonth->" + err.Error())
		}
		runtime.Gosched()
	}()

	go func() {
		defer wg.Done()

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		err := sensor2_TableMonth(*ticker, ctx, monthRepo)
		if err != nil {
			log.Fatalf("Error sensor2_TableMonth->" + err.Error())
		}

		runtime.Gosched()
	}()

	s := <-stoped
	switch s {
	case syscall.SIGHUP:
		log.Println("[hungup]")
	case syscall.SIGINT:
		log.Println("[interupt]")
	case syscall.SIGTERM:
		log.Println("[force stop]")
	case syscall.SIGQUIT:
		log.Println("[stop and core dump]")
	default:
		log.Println("[Unknown signal]")
	}

}

func sensor1_TableMonth(ticker time.Ticker, ctx context.Context, monthRepo month.IMonthRepository) error {
	var settingan models.Settings
	temp := float32(15.15)
	rh := float32(65.15)
	for {
		select {
		case <-ticker.C:
			log.Println("[sensor 1]")
			var data models.Month
			data.Dev_ID = settingan.Dev_ID
			data.Sensor_ID = 1
			ts := time.Now()
			data.Timestamp = ts.Format(time.RFC3339)
			data.Temp = temp
			data.Rh = rh

			temp += 1.5
			rh += 1.5
			if temp > 65 {
				temp = 15.15
			}
			if rh >= 100 {
				rh = 65.15
			}

			month := fmt.Sprintf("%02d", ts.Month())
			err := monthRepo.InsertDatatoTable(ctx, month, data)
			if err != nil {
				log.Fatalf("error InsertMonth2 :%v", err)
			}

		case <-ctx.Done():
			return ctx.Err()

		}
	}
}

func sensor2_TableMonth(ticker time.Ticker, ctx context.Context, monthRepo month.IMonthRepository) error {
	var settingan models.Settings
	temp := float32(25.15)
	rh := float32(75.15)
	for {
		select {
		case <-ticker.C:
			log.Println("======== [sensor 2]")
			var data models.Month
			data.Dev_ID = settingan.Dev_ID
			data.Sensor_ID = 2
			ts := time.Now()
			data.Timestamp = ts.Format(time.RFC3339)
			data.Temp = temp
			data.Rh = rh

			temp += 1.5
			rh += 1.5
			if temp > 65 {
				temp = 15.15
			}
			if rh >= 100 {
				rh = 65.15
			}

			date := fmt.Sprintf("%02d", ts.Month())
			err := monthRepo.InsertDatatoTable(ctx, date, data)
			if err != nil {
				log.Fatalf("error InsertMonth2 :%v", err)
			}

		case <-ctx.Done():
			return ctx.Err()

		}
	}
}
