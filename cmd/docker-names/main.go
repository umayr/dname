package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"

	dn "github.com/umayr/docker-names"
)

func main() {
	app := cli.NewApp()

	app.Version = "0.0.1"
	app.Name = "docker-names"
	app.Usage = "generate names like docker does when it creates a new container"
	app.UsageText = "docker-names [command] [arguments...]"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name: "number, n",
			Value: 1,
			Usage: "total number of generated names",
		},
	}
	app.Action = func(c *cli.Context) error {
		for i := 0; i < c.Int("number"); i++ {
			fmt.Println(dn.Generate())
		}
		return nil
	}

	app.Run(os.Args)
}
