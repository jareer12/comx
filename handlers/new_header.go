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
			Label:        "Select the compiler you want to use",
			Items:        header_exts,
			HideSelected: true,
		}

		_, extension, err := prompt.Run()

		if err != nil {
			utils.PrintError("You did not select a proper compiler")
			return nil
		}

		os.WriteFile(
			fmt.Sprintf("%v/%v/%v.%v", cwd, store.MainDir, file_name, extension),
			[]byte(fmt.Sprintf("#ifndef HEADER_%v\n\n#define HEADER_%v\n// ---snip---\n#endif",
				module_name,
				module_name),
			),
			os.ModePerm,
		)
	} else {
		utils.PrintError("Enter header file name, usage: 'example header <name>'")
	}

	return nil
}
