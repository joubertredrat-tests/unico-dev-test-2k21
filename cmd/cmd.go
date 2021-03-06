package cmd

import (
	"log"
	"os"

	"github.com/joubertredrat-tests/unico-dev-test-2k21/infra/api"
	"github.com/joubertredrat-tests/unico-dev-test-2k21/infra/worker"
	"github.com/urfave/cli/v2"
)

func Execute() {
	app := &cli.App{
		Name:     "Unico dev test",
		Usage:    "App from unico tech dev test",
		HideHelp: true,
		Commands: []*cli.Command{
			{
				Name:    "api",
				Aliases: []string{},
				Usage:   "Start listen http API",
				Action: func(c *cli.Context) error {
					return api.Run()
				},
			},
			{
				Name:    "import",
				Aliases: []string{},
				Usage:   "Import open market list from file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "filename",
						Aliases:  []string{},
						Usage:    "Filename path to import open market list",
						Value:    "",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					return worker.RunImport(c.String("filename"))
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
