package actions

import (
	"log"

	working "github.com/caarlos0/am-i-working"
	"github.com/urfave/cli"
)

// Main action of the app
func Main(c *cli.Context) error {
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
}
