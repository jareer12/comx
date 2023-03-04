package handlers

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	DefaultHeadCode = "#include <stdio.h>\n\nint main() {\n	return 0;\n}"
)

func NewHeaderHandler(cCtx *cli.Context) error {
	file_name := cCtx.Args().Get(0)

	if len(file_name) > 0 {
		os.WriteFile(fmt.Sprintf("./src/%v", file_name), []byte(DefaultHeadCode), os.ModePerm)
	} else {

	}

	return nil
}
