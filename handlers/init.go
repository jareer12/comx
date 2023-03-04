package handlers

import (
	"fmt"
	"os"
	"strings"
	"time"

	"cli/utils"

	"github.com/urfave/cli/v2"
)

const (
	DefaultCode    = "#include <stdio.h>\n\nint main() {\n  // ---snip---\n	return 0;\n}"
	DefaultCppCode = "#include <iostream>\n\nint main() {\n  // ---snip---\n  return 0;\n}"
)

func CreateMain(file_ext string, def_code string, proj_lang string) error {
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
		Compilers:   []string{},
		ProjectLang: proj_lang,
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

func ProjectExists() bool {
	// Check if a project already exists in the target dir
	// Check for config file

	if _, err := os.ReadFile(utils.StorePath); err == nil {
		return true
	} else {
		return false
	}
}

func InitHandle(cCtx *cli.Context) error {
	proj_lang := cCtx.Args().Get(0)

	if ProjectExists() {
		config, err := utils.StoreContents()

		if err != nil {
			utils.PrintError("Prohect seems to be corrupted")
			return nil
		}

		utils.PrintError(fmt.Sprintf("A '%v' project already exists in the mentioned directory.", strings.ToUpper(config.ProjectLang)))
		return nil
	}

	switch proj_lang {
	case "c":
		{
			CreateMain(proj_lang, DefaultCode, "C")
		}
	case "cpp":
		{
			start := time.Now().UnixMicro()
			CreateMain(proj_lang, DefaultCppCode, "C++")
			utils.PrintSuccess(fmt.Sprintf("Successfuly created new C++ project, elapsed %vms.", (time.Now().UnixMicro()-start)/1000))
		}
	default:
		{
			utils.PrintError(fmt.Sprintf("Selected language '%v' not supported, valid options are 'c' or 'cpp'\n", proj_lang))
		}
	}

	FindCompilersMain(false)

	return nil
}
