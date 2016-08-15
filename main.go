package main

import (
	"os"

	"github.com/caarlos0/am-i-working/watcher"
	"github.com/urfave/cli"
)

var version = "dev"

func main() {
	app := cli.NewApp()
	app.Name = "am-i-working"
	app.Usage = "Logs when you're working based on /etc/resolv.conf domain"
	app.Version = version
	app.Author = "Carlos Alexandro Becker <@caarlos0>"
	app.Copyright = "MIT"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "domain, d",
			Usage: "Domain name that appears in domain section of /etc/resolv.conf when you're connected to work networks",
		},
	}
	app.Action = func(c *cli.Context) error {
		if err := watcher.Watch(c.String("domain")); err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
		return nil
	}
	app.Run(os.Args)
}
