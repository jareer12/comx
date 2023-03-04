package handlers

import (
	"cli/utils"
	"fmt"
	"log"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func BuildHandler(cCtx *cli.Context) error {
	if store, err := utils.StoreContents(); err != nil {
		return nil
	} else {
		var main_file string

		fn := cCtx.Args().Get(0)

		if len(fn) > 0 {
			main_file = fn
		} else {
			main_file = fmt.Sprintf("main.%v", store.ProjectName)
		}

		compiler := store.SelectedCompiler
		cmd := exec.Command(fmt.Sprintf("%v ./src/%v", compiler, main_file))

		err := cmd.Run()

		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
