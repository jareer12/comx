package handlers

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

const (
	version = "1.0.0"
)

func VersionHandle(cCtx *cli.Context) error {
	fmt.Printf("Version: v%v\n", version)
	return nil
}
