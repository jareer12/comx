package handlers

import (
	"cli/utils"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/urfave/cli/v2"
)

var (
	ex_compilers = []string{"gcc", "g++", "cpp"}
)

func FindCompilersHandle(cCtx *cli.Context) error {
	envs := os.Environ()
	os_name := runtime.GOOS

	for i := 0; i < len(ex_compilers); i++ {
		for j := 0; j < len(envs); j++ {
			if strings.Contains(envs[j], ex_compilers[i]) {
				utils.PrintInfo(fmt.Sprintf("Found compiler in env variables: %v", envs[j]))
			}
		}

		if os_name == "linux" {
			dirs, err := os.ReadDir("/usr/bin")

			if err != nil {
				utils.PrintError("Unable to read /usr/bin directory.")
			}

			for k := 0; k < len(dirs); k++ {
				if dirs[k].Name() == ex_compilers[i] {
					utils.PrintInfo(fmt.Sprintf("Found compiler in bin directory: /usr/bin/%v", dirs[k].Name()))
				}
			}
		}
	}

	return nil
}
