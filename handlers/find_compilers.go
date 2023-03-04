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

func AddCompiler(path string, ys utils.YamlStore) utils.YamlStore {
	ys.Compilers = append(ys.Compilers, path)
	return ys
}

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
					comp_path := fmt.Sprintf("/usr/bin/%v", dirs[k].Name())
					if config, err := os.ReadFile("./config.yaml"); err == nil {
						ys, d_err := utils.DecodeStore(config)

						if d_err != nil {
							utils.PrintError("ff")
						}

						fmt.Println(ys)

						// new := AddCompiler(comp_path, ys)
						// utils.SaveConfig(new)

						utils.PrintInfo(fmt.Sprintf("Found compiler in bin directory: %v", comp_path))

					} else {
						utils.PrintError(fmt.Sprintf("Unable to add compiler(%v) to config", comp_path))
					}
				}
			}
		}
	}

	return nil
}
