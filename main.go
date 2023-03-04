package main

import (
	"log"
	"os"

	"cli/handlers"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "ComX",
		Usage: "Manage your C/C++ projects with a CLI tool.",
		Commands: []*cli.Command{
			{
				Name:    "initialize",
				Aliases: []string{"init", "new"},
				Usage:   "Initialize a new C/C++ project.",
				Action:  handlers.InitHandle,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "main_file",
						Aliases: []string{"mf"},
						Usage:   "Name of the file inside the main code directory.",
					},
					&cli.StringFlag{
						Name:    "main_dir",
						Aliases: []string{"md"},
						Usage:   "Name of the main code directory.",
					},
				},
			},
			{
				Name:    "build",
				Aliases: []string{"b"},
				Action:  handlers.BuildHandler,
				Usage:   "Build the project using a pre-selected compiler",
			},
			{
				Name:    "header",
				Aliases: []string{"h"},
				Action:  handlers.NewHeaderHandler,
				Usage:   "Create a new header file for a project.",
			},
			{
				Aliases: []string{"fc"},
				Name:    "find-compilers",
				Action:  handlers.FindCompilersHandle,
				Usage:   "Search your machine for possible GNU compilers.",
			},
			{
				Aliases: []string{"sc"},
				Name:    "select-compiler",
				Action:  handlers.SelectCompilerHandle,
				Usage:   "Select the founded compilers.",
			},
			{
				Name:    "list-compilers",
				Action:  handlers.ListCompilers,
				Aliases: []string{"compilers", "lc"},
				Usage:   "List all founded compilers.",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
