package handlers

import (
	"fmt"
	"os"
	"strings"
	"time"

	"cli/utils"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

const (
	DefaultCode    = "#include <stdio.h>\n\nint main() {\n  // ---snip---\n	return 0;\n}"
	DefaultCppCode = "#include <iostream>\n\nint main() {\n  // ---snip---\n  return 0;\n}"
)

func CreateMain(file_ext string, proj_name string, def_code string, proj_lang string, main_file string, main_dir string, mod_dir string) error {
	cwd, cerr := os.Getwd()

	if cerr != nil {
		return nil
	}

	err := os.Mkdir(fmt.Sprintf("%v/%v", cwd, main_dir), os.ModePerm)

	if err != nil {
		return err
	}

	cd_err := os.Mkdir(fmt.Sprintf("%v/config", cwd), os.ModePerm)

	if cd_err != nil {
		return cd_err
	}

	f_err := os.WriteFile(fmt.Sprintf("%v/%s/%s", cwd, main_dir, main_file), []byte(def_code), os.ModePerm)

	if f_err != nil {
		return f_err
	}

	def_conf, c_err := utils.StoreToText(utils.StoreStruct{
		Compilers:   []string{},
		ProjectLang: proj_lang,
		ProjectName: proj_name,
		MainFile:    main_file,
		MainDir:     main_dir,
		ModulesDir:  mod_dir,
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
	if _, err := os.ReadFile(utils.StorePath); err == nil {
		return true
	} else {
		return false
	}
}

func InitHandle(cCtx *cli.Context) error {
	if ProjectExists() {
		config, err := utils.StoreContents()

		if err != nil {
			utils.PrintError("Prohect seems to be corrupted")
			return nil
		}

		utils.PrintError(fmt.Sprintf("A '%v' project already exists in the mentioned directory.", strings.ToUpper(config.ProjectLang)))
		return nil
	}

	var main_file string
	var mod_dir string
	var main_dir string

	proj_name := cCtx.Args().Get(0)

	if len(proj_name) == 0 || proj_name == " " {
		utils.PrintError("Provide a valid project name 'example init <project-name>'")
		return nil
	}

	prompt := promptui.Select{
		Label:        "Please select project base language",
		Items:        []string{"cpp", "c"},
		HideSelected: true,
	}

	_, proj_lang, err := prompt.Run()

	if err != nil {
		utils.PrintError("You did not select a project language properly")
		return nil
	}

	mf := cCtx.String("main_file")
	mod := cCtx.String("mod_dir")
	md := cCtx.String("main_dir")

	if len(mf) > 0 {
		main_file = mf
	} else {
		main_file = fmt.Sprintf("main.%v", proj_lang)
	}

	if len(md) > 0 {
		main_dir = md
	} else {
		main_dir = "src"
	}

	if len(mod) > 0 {
		mod_dir = mod
	} else {
		mod_dir = "modules"
	}

	switch proj_lang {
	case "c":
		{
			CreateMain(proj_lang, proj_name, DefaultCode, "C", main_file, main_dir, mod_dir)
		}
	case "cpp":
		{
			start := time.Now().UnixMicro()
			CreateMain(proj_lang, proj_name, DefaultCppCode, "C++", main_file, main_dir, mod_dir)

			elapsed := float64(time.Now().UnixMicro()-start) / 1000
			utils.PrintSuccess(fmt.Sprintf("Successfuly created new C++ project, elapsed %vms.", elapsed))
		}
	default:
		{
			utils.PrintError(fmt.Sprintf("Selected language '%v' not supported, valid options are 'c' or 'cpp'\n", proj_lang))
		}
	}

	FindCompilersMain(false)

	return nil
}
