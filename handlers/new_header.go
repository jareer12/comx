package handlers

import (
	"cli/utils"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

var (
	header_exts = []string{"h", "hh", "cc", "hpp", "hxx"}
)

func AddIncludeToFile(file_path string, include_path string) error {
	if bytes, err := os.ReadFile(file_path); err != nil {
		return err
	} else {
		or := fmt.Sprintf("#include <%v>\n%v", include_path, string(bytes))
		err := os.WriteFile(file_path, []byte(or), os.ModePerm)

		if err != nil {
			return err
		}
	}

	return nil
}

func NewHeaderHandler(cCtx *cli.Context) error {
	store, s_err := utils.StoreContents()

	if s_err != nil {
		return s_err
	}

	cwd, cw_err := os.Getwd()

	if cw_err != nil {
		return cw_err
	}

	file_name := cCtx.Args().Get(0)
	if len(file_name) > 0 {
		module_name := strings.ToUpper(strings.Split(file_name, ".")[0])

		prompt := promptui.Select{
			Label:        "Select the module extension.",
			Items:        header_exts,
			HideSelected: true,
		}

		_, extension, err := prompt.Run()

		if err != nil {
			utils.PrintError("You did not select a proper module extension")
			return nil
		}

		os.WriteFile(
			fmt.Sprintf("%v/%v/%v.%v", cwd, store.ModulesDir, file_name, extension),
			[]byte(fmt.Sprintf("#ifndef HEADER_%v\n\n#define HEADER_%v\n// ---snip---\n#endif",
				module_name,
				module_name),
			),
			os.ModePerm,
		)

		module_path := fmt.Sprintf("../%v/%v.%v", store.ModulesDir, file_name, extension)
		include_err := AddIncludeToFile(fmt.Sprintf("%v/%v/%v", cwd, store.MainDir, store.MainFile), module_path)

		if include_err != nil {
			utils.PrintError("Unable to add module include to main file.")
			println(include_err)
		}
	} else {
		utils.PrintError("Enter header file name, usage: 'example header <name>'")
	}

	return nil
}
