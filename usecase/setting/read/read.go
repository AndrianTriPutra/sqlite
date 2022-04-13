package read

import (
	"context"
	"fmt"
	"log"
	"xsqlitex/pkg/repository/setting"
)

func Read(ctx context.Context, settingRepo setting.ISettingRepository) {
	data, err := settingRepo.ReadAll(ctx)
	if err != nil {
		log.Fatalf("error read all setting : %s", err)
	}

	for _, val := range data {
		fmt.Print(val.ID)
		fmt.Print(" | ")
		fmt.Print(val.Parameter)
		fmt.Print(" | ")
		fmt.Println(val.Value)
	}
}
