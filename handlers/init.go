package handlers

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	DefaultCode = "#include<stdio.h>\nint main() {\n}"
)

func CreateMain(file_ext string) error {
	err := os.Mkdir("./src", os.ModePerm)

	if err != nil {
		return err
	} else {
		err := os.WriteFile(fmt.Sprintf("./src/main.%v", file_ext), []byte(DefaultCode), os.ModePerm)

		if err != nil {
			return err
		}
	}

	return nil
}

func InitHandle(cCtx *cli.Context) error {
	proj_lang := cCtx.String("lang")

	switch proj_lang {
	case "c":
		{
			fmt.Println("Creating new C project, please wait.")
			CreateMain(proj_lang)
		}
	case "cpp":
		{
			fmt.Println("Creating new C++ project, please wait.")
			CreateMain(proj_lang)
		}
	default:
		{
			fmt.Print(fmt.Sprintf("Selected language '%v' not supported, valid options are: 'c' or 'cpp'", proj_lang))
		}
	}

	return nil
}
