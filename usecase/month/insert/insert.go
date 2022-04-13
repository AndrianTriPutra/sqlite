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

var Setting models.Settings

func Insert(ctx context.Context, monthRepo month.IMonthRepository, settingRepo setting.ISettingRepository) {
	//log.Println(" insert 2 value to 1 table")

	data, err := settingRepo.ReadAll(ctx)
	if err != nil {
		log.Fatalf("error read setting : %s", err)
	}

	for _, val := range data {
		switch val.Parameter {
		case "device_id":
			Setting.Dev_ID = val.Value
		case "interval":
			Setting.Interval = val.Value
		}
	}
	intInterval, _ := strconv.Atoi(Setting.Interval)
	interval := time.Duration(intInterval)
	interval = interval * time.Minute
	log.Printf("device_id:%s", Setting.Dev_ID)
	log.Printf("interval :%v", interval)

	var wg sync.WaitGroup
	stopchan := make(chan bool, 1)
	kill := make(chan os.Signal, 1)
	signal.Notify(kill,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		wg.Wait()

		close(stopchan)
	}()

	wg.Add(2)

	go func() {
		defer wg.Done()

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		sensor1_TableMonth(*ticker, stopchan, ctx, monthRepo)
		runtime.Gosched()
	}()

	go func() {
		defer wg.Done()

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		sensor2_TableMonth(*ticker, stopchan, ctx, monthRepo)
		runtime.Gosched()
	}()

	exit_chan := make(chan int)
	go func() {
		for {
			s := <-kill
			switch s {
			// kill -SIGHUP XXXX
			case syscall.SIGHUP:
				log.Println("close_cause [hungup]")
				exit_chan <- 3

			// kill -SIGINT XXXX or Ctrl+c
			case syscall.SIGINT:
				log.Println("close_cause [interupt]")
				exit_chan <- 2

			// kill -SIGTERM XXXX
			case syscall.SIGTERM:
				log.Println("close_cause [force_stop]")
				exit_chan <- 0

			// kill -SIGQUIT XXXX
			case syscall.SIGQUIT:
				log.Println("close_cause [stop and core dump]")
				exit_chan <- 0

			default:
				log.Println("close_cause [Unknown signal]")
				exit_chan <- 1
			}
		}
	}()

	code := <-exit_chan
	log.Println("< close one >")
	os.Exit(code)
}

func sensor1_TableMonth(ticker time.Ticker, stop chan bool, ctx context.Context, monthRepo month.IMonthRepository) {
	temp := float32(15.15)
	rh := float32(65.15)
	for {
		select {
		case <-ticker.C:
			log.Println("[sensor 1]")
			var data models.Month
			data.Dev_ID = Setting.Dev_ID
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

		case <-stop:
			log.Println("stop sensor1_TableMonth")
			os.Exit(1)
		}
	}
}

func sensor2_TableMonth(ticker time.Ticker, stop chan bool, ctx context.Context, monthRepo month.IMonthRepository) {
	temp := float32(25.15)
	rh := float32(75.15)
	for {
		select {
		case <-ticker.C:
			log.Println("======== [sensor 2]")
			var data models.Month
			data.Dev_ID = Setting.Dev_ID
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

		case <-stop:
			log.Println("stop sensor2_TableMonth")
			os.Exit(1)
		}
	}
}
