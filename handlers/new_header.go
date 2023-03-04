package handlers

import (
	"cli/utils"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func NewHeaderHandler(cCtx *cli.Context) error {
	file_name := cCtx.Args().Get(0)
	if len(file_name) > 0 {
		module_name := strings.ToUpper(strings.Split(file_name, ".")[0])
		os.WriteFile(
			fmt.Sprintf("./src/%v.h", file_name),
			[]byte(fmt.Sprintf("#ifndef %v\n\n#define %v\n// ---snip---\n#endif",
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
