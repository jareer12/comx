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

	f_err := os.WriteFile(fmt.Sprintf("./src/main.%v", file_ext), []byte(def_code), os.ModePerm)

	if f_err != nil {
		return f_err
	}

	def_conf, c_err := utils.EncodeStore(utils.YamlStore{
		Compilers: []string{},
	})

	if c_err != nil {
		return c_err
	}

	fc_err := os.WriteFile("./config.yaml", []byte(def_conf), os.ModePerm)

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
			fmt.Println("Creating new C project, please wait.")
			CreateMain(proj_lang, DefaultCode)
		}
	case "cpp":
		{
			start := time.Now().UnixMicro()
			utils.PrintInfo("Creating a new C++ project, please wait.")

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
