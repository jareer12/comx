package handlers

import (
	"fmt"
	"os"
	"time"

	"cli/utils"

	"github.com/urfave/cli/v2"
)

const (
	DefaultCode    = "#include <stdio.h>\n\nint main() {\n  // ---snip---\n	return 0;\n}"
	DefaultCppCode = "#include <iostream>\n\nint main() {\n  // ---snip---\n  return 0;\n}"
)

func CreateMain(file_ext string, def_code string) error {
	err := os.Mkdir("./src", os.ModePerm)

	if err != nil {
		return err
	}

	cd_err := os.Mkdir("./config", os.ModePerm)

	if cd_err != nil {
		return cd_err
	}

	f_err := os.WriteFile(fmt.Sprintf("./src/main.%v", file_ext), []byte(def_code), os.ModePerm)

	if f_err != nil {
		return f_err
	}

	def_conf, c_err := utils.StoreToText(utils.StoreStruct{
		Compilers: []string{},
	})

	if c_err != nil {
		return c_err
	}

	fc_err := os.WriteFile(utils.StorePath, []byte(def_conf), os.ModePerm)

	if fc_err != nil {
		return fc_err
	}

	return nil
}

func InitHandle(cCtx *cli.Context) error {
	proj_lang := cCtx.Args().Get(0)

	switch proj_lang {
	case "c":
		{
			CreateMain(proj_lang, DefaultCode)
		}
	case "cpp":
		{
			start := time.Now().UnixMicro()
			CreateMain(proj_lang, DefaultCppCode)
			utils.PrintSuccess(fmt.Sprintf("Successfuly created project, elapsed %vms.", (time.Now().UnixMicro()-start)/1000))
		}
	default:
		{
			utils.PrintError(fmt.Sprintf("Selected language '%v' not supported, valid options are 'c' or 'cpp'\n", proj_lang))
		}
	}

	return nil
}
