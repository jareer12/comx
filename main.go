package main

import (
	"log"
	"os"

	"cli/handlers"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "CC-CLI",
		Usage: "Manage your C/C++ projects with a CLI tool.",
		// Action: func(*cli.Context) error {
		// 	return nil
		// },
		Commands: []*cli.Command{
			{
				Name:    "initialize",
				Aliases: []string{"init"},
				Usage:   "Setup a new C/C++ project.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "lang",
						Usage:    "Select project language, '--lang=c' or '--lang=cpp'.",
					},
				},
				Action: handlers.InitHandle,
			},
			{
				Name:    "header",
				Aliases: []string{"h"},
				Usage:   "Create a new header file for a project.",
				Flags: []cli.Flag{},
				Action: handlers.NewHeaderHandler,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
