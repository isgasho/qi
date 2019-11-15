package main

import (
	"os"

	"github.com/mattn/qi/cmd"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "qiitasearch"
	app.Usage = "search qiita articles"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name:  "myqi",
			Usage: "qiita + mine : you get yours articles",
			Action: func(c *cli.Context) error {
				qiitaToken := os.Getenv("QIITA_TOKEN")
				data := cmd.FetchQiitaData(qiitaToken)
				cmd.OutputQiitaData(data)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
