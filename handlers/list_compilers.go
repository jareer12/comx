package handlers

import (
	"cli/utils"
	"fmt"

	"github.com/urfave/cli/v2"
)

func ListCompilers(cCtx *cli.Context) error {
	if store, err := utils.StoreContents(); err != nil {
		fmt.Println(err)
		utils.PrintError(fmt.Sprintf("Error, unable to parse config file(%v), maybe it's corrupted.", utils.StorePath))
	} else {
		len := len(store.Compilers)

		if len == 0 {
			utils.PrintInfo("No compilers found, use the 'find-compilers' command to find some.")
			return nil
		}

		for i := 0; i < len; i++ {
			utils.PrintInfo(store.Compilers[i])
		}

		utils.PrintInfo(fmt.Sprintf("Total %v compilers found", len))
	}

	return nil
}
