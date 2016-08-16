package main

import (
	"log"
	"os"

	"github.com/caarlos0/am-i-working/cmd/actions"
	"github.com/urfave/cli"
)

var version = "dev"

func main() {
	app := cli.NewApp()
	app.Name = "am-i-working"
	app.Usage = "Logs your working activity based on /etc/resolv.conf domain"
	app.Version = version
	app.Author = "Carlos Alexandro Becker <@caarlos0>"
	app.Copyright = "MIT"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "domain, d",
			Usage: "Domain name that appears in domain section of /etc/resolv.conf when you're connected to company networks",
		},
		cli.StringFlag{
			Name:  "file, f",
			Usage: "File to watch for domain regexes",
			Value: "/etc/resolv.conf",
		},
	}
	app.Action = actions.Main
	log.Fatal(app.Run(os.Args))
}
