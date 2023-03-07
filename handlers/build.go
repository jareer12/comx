package handlers

import (
	"cli/utils"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"
)

func BuildHandler(cCtx *cli.Context) error {
	if store, err := utils.StoreContents(); err != nil {
		return nil
	} else {
		var main_file string
		var file_output string

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

		cwd, c_err := os.Getwd()

		if c_err != nil {
			utils.PrintError("Something went wrong while getting the CWD path, try running as root")
			return nil
		}

		output := cCtx.String("output")

		if len(output) == 0 || output == " " {
			utils.PrintError("Please enter a valid output path using the '--output' argument.")
			return nil
		}

		if strings.HasPrefix(output, "./") {
			re, _ := regexp.Compile("/+")
			file_output = re.ReplaceAllLiteralString(fmt.Sprintf("%v/%v", cwd, strings.Trim(output, "./")), "/")
		} else {
			file_output = output
		}

		compiler := store.SelectedCompiler

		if !(len(compiler) > 0) {
			utils.PrintError("Please select a compiler using the 'select-compiler' command")
			return nil
		}

		build_path := fmt.Sprintf("%v/%s/%s", cwd, main_dir, main_file)

		utils.PrintInfo(fmt.Sprintf(`Building using command: %v %v -o %v`, compiler, build_path, file_output))
		utils.PrintInfo("If build doesn't succeed, select a different compiler using the 'select-compiler' command.")

		cmd := exec.Command(compiler, build_path, "-o", file_output)
		outp, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + string(outp))
			return nil
		} else {
			println(string(outp))
		}

		utils.PrintSuccess("Build Succeeded")
	}

	return nil
}
