package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"xsqlitex/pkg/repository"
	"xsqlitex/pkg/sqLite"
	m_convert "xsqlitex/usecase/month/convert"
	m_delete "xsqlitex/usecase/month/delete"
	m_insert "xsqlitex/usecase/month/insert"
	m_read "xsqlitex/usecase/month/read"
	s_insert "xsqlitex/usecase/setting/insert"
	s_read "xsqlitex/usecase/setting/read"
	s_update "xsqlitex/usecase/setting/update"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	db := sqLite.NewsqliteDb("example.db")
	err := db.Connect(ctx)
	if err != nil {
		log.Fatalf("cannot connect to database :%v", err)
		return
	}

	monthRepo, settingRepo := repository.Init(db)

	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		fmt.Printf("     go run . s_insert\n")
		fmt.Printf("     go run . s_read\n")
		fmt.Printf("     go run . s_update\n")
		fmt.Printf("     go run . m_insert\n")
		fmt.Printf("     go run . m_read\n")
		fmt.Printf("     go run . m_delete\n")
		fmt.Printf("     go run . m_csv\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	log.Printf("methode:%s", flag.Args()[0])

	switch flag.Args()[0] {
	case "s_insert":
		fmt.Println()
		s_insert.Insert(ctx, settingRepo)

	case "s_read":
		fmt.Println()
		s_read.Read(ctx, settingRepo)

	case "s_update":
		fmt.Println()
		s_update.Update(ctx, settingRepo)

	case "m_insert":
		fmt.Println()
		m_insert.Insert(ctx, cancel, monthRepo, settingRepo)

	case "m_read":
		fmt.Println()
		m_read.Read(ctx, monthRepo, settingRepo)

	case "m_delete":
		fmt.Println()
		m_delete.Deleted(ctx, monthRepo, settingRepo)

	case "m_csv":
		fmt.Println()
		m_convert.ExportoCSV(ctx, monthRepo, settingRepo)

	}

}
