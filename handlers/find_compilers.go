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
	ex_compilers   = []string{"gcc", "g++", "cpp", "cc", "c++", "cpp-11", "clang"}
	win_comps_locs = []string{"mingw"}
)

func AddCompiler(path string, ys utils.StoreStruct) utils.StoreStruct {
	ys.Compilers = append(ys.Compilers, path)

	return ys
}

func MatchWinCompilers(path string) bool {
	for i := 0; i < len(win_comps_locs); i++ {
		if strings.Contains(path, win_comps_locs[i]) {
			return true
		}
	}

	return false
}

func FindCompilersMain(verbose bool) error {
	envs := os.Environ()
	os_name := runtime.GOOS

	for i := 0; i < len(ex_compilers); i++ {
		for j := 0; j < len(envs); j++ {
			if os_name == "win" || os_name == "windows" {
				if MatchWinCompilers(envs[j]) {
					path := envs[j]
					if read_bin, err := os.ReadDir(path); err != nil {
						utils.PrintError(fmt.Sprintf("Error reading the following directory: %v", path))
						fmt.Println(err)
					} else {
						for r := 0; r < len(read_bin); r++ {
							pc := read_bin[r]

							for e := 0; e < len(ex_compilers); e++ {
								ec := ex_compilers[e]
								pn := pc.Name()

								if strings.Contains(pn, ec) && !pc.IsDir() && strings.HasSuffix(pn, "exe") {
									utils.PrintInfo(fmt.Sprintf("Found compiler from: %v, compiler: %v", path, pc))
								}
							}
						}
					}
				}
			}

			if strings.Contains(envs[j], ex_compilers[i]) {
				if verbose {
					utils.PrintInfo(fmt.Sprintf("Found compiler in env variables: %v", envs[j]))
				}
			}
		}

		if os_name == "linux" {
			dirs, err := os.ReadDir("/usr/bin")

			if err != nil {
				if verbose {
					utils.PrintError("Unable to read /usr/bin directory.")
				}
			}

			for k := 0; k < len(dirs); k++ {
				dir := dirs[k]

				if dir.Name() == ex_compilers[i] && !dir.IsDir() {
					comp_path := fmt.Sprintf("/usr/bin/%v", dir.Name())
					if verbose {
						utils.PrintInfo(fmt.Sprintf("Found compiler in bin directory: %v", comp_path))
					}

					if cbytes, err := os.ReadFile(utils.StorePath); err == nil {
						if store, err := utils.GetStore(cbytes); err != nil {
							if verbose {
								utils.PrintError(fmt.Sprintf("Something went wrong, maybe config file(%v) is corrupted.", utils.StorePath))
							}
						} else {
							if !utils.HasCompiler(comp_path, store) {
								utils.AddCompiler(comp_path, &store)
							}

							utils.SaveConfig(store)
						}
					} else {
						if verbose {
							utils.PrintError(fmt.Sprintf("Unable to add compiler(%v) to config", comp_path))
						}
					}
				}
			}
		}
	}

	return nil
}
func FindCompilersHandle(_ *cli.Context) error {
	return FindCompilersMain(true)
}
