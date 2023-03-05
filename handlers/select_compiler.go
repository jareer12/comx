package handlers

import (
	"cli/utils"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

func SelectCompilerHandle(cCtx *cli.Context) error {
	store, err := utils.StoreContents()
	if err != nil {
		return nil
	}

	prompt := promptui.Select{
		Label:        "Select the compiler you want to use",
		Items:        store.Compilers,
		HideSelected: true,
	}

	_, compiler, err := prompt.Run()

	if err != nil {
		utils.PrintError("You did not select a proper compiler")
		return nil
	}

	utils.SetCompiler(compiler, &store)
	utils.SaveConfig(store)
	utils.PrintSuccess(fmt.Sprintf("Selected compiler: %v", compiler))

	return nil
}
