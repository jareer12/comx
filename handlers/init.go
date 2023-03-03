package handlers

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"

	"github.com/urfave/cli/v2"
)

const (
	DefaultCode    = "#include <stdio.h>\n\nint main() {\n	return 0;\n}"
	DefaultCppCode = "#include <iostream>\n\nint main() {\n	 return 0;\n}"
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

	return nil
}

func InitHandle(cCtx *cli.Context) error {
	proj_lang := cCtx.String("lang")

	switch proj_lang {
	case "c":
		{
			fmt.Println("Creating new C project, please wait.")
			CreateMain(proj_lang, DefaultCode)
		}
	case "cpp":
		{
			start := time.Now().UnixMicro()
			fmt.Printf(fmt.Sprintf("%v Creating a new C++ project, please wait.\n", color.CyanString("►")))

			CreateMain(proj_lang, DefaultCppCode)
			fmt.Printf(fmt.Sprintf("%v Successfuly created new project, elapsed %vms", color.GreenString("✔"), (time.Now().UnixMicro()-start)/1000))
		}
	default:
		{
			fmt.Printf(fmt.Sprintf("Selected language '%v' not supported, valid options are: 'c' or 'cpp'\n", proj_lang))
		}
	}

	return nil
}
