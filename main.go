package main

import (
	"log"
	"os"

	"github.com/caarlos0/am-i-working/watcher"
	"github.com/urfave/cli"
)

var version = "dev"

const resolv = "/etc/resolv.conf"

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
		events := make(chan bool)
		domain := c.String("domain")
		log.Println("Watching for domain", domain)
		go func() {
			if err := watcher.Watch(resolv, domain, events); err != nil {
				log.Fatalln(err)
			}
		}()
		for {
			if <-events {
				log.Println("Working")
			} else {
				log.Println("Not working")
			}
		}
		return nil
	}
	app.Run(os.Args)
}
