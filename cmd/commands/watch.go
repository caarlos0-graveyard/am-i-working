package commands

import (
	"log"

	working "github.com/caarlos0/am-i-working"
	"github.com/urfave/cli"
)

// Watch action of the app
var Watch = cli.Command{
	Name:  "watch",
	Usage: "Watch for domain changes",
	Flags: flags,
	Action: func(c *cli.Context) error {
		events := make(chan bool)
		domain := c.String("domain")
		resolv := c.String("file")
		log.Println("Watching", resolv, "for domain", domain)
		go func() {
			if err := working.Watch(resolv, domain, events); err != nil {
				log.Fatal(err)
			}
		}()
		for {
			if <-events {
				log.Println("Working")
			} else {
				log.Println("Not working")
			}
		}
	},
}
