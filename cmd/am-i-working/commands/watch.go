package commands

import (
	"log"

	amiworking "github.com/caarlos0/am-i-working"
	"github.com/urfave/cli"
)

// Watch action of the app
var Watch = cli.Command{
	Name:  "watch",
	Usage: "Watch for domain changes",
	Flags: flags,
	Action: func(c *cli.Context) error {
		var events = make(chan bool)
		var domain = c.String("domain")
		var resolv = c.String("file")
		if domain == "" {
			return cli.NewExitError("missing domain name", 1)
		}
		log.Println("Watching", resolv, "for domain", domain)
		go func() {
			if err := amiworking.Watch(resolv, domain, events); err != nil {
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
