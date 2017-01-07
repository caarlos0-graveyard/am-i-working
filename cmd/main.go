package main

import (
	"log"
	"os"

	"github.com/caarlos0/am-i-working/cmd/commands"
	"github.com/urfave/cli"
)

var version = "master"

func main() {
	app := cli.NewApp()
	app.Name = "am-i-working"
	app.Usage = "Logs your working activity based on /etc/resolv.conf domain"
	app.Version = version
	app.Author = "Carlos Alexandro Becker <@caarlos0>"
	app.Copyright = "MIT"
	app.Commands = []cli.Command{
		commands.Watch,
		commands.CreateService,
		commands.RestartService,
		commands.StartService,
		commands.StopService,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
