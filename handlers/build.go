package handlers

import (
	"cli/utils"
	"fmt"
	"log"
	"os"
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
			main_file = store.MainFile
		}

		var main_dir string

		dn := cCtx.Args().Get(0)

		if len(dn) > 0 {
			main_dir = dn
		} else {
			main_dir = store.MainDir
		}

		compiler := store.SelectedCompiler

		if !(len(compiler) > 0) {
			utils.PrintError("Please select a compiler using the 'select-compiler' command")
			return nil
		}

		cwd, c_err := os.Getwd()

		if c_err != nil {
			utils.PrintError("Something went wrong while getting the CWD path, try running as root")
			return nil
		}

		build_path := fmt.Sprintf("%v/%s/%s", cwd, main_dir, main_file)
		utils.PrintSuccess(fmt.Sprintf("Building using command: %v %v", compiler, build_path))

		cmd := exec.Command(compiler, build_path)
		err := cmd.Run()

		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
