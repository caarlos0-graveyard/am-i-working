package main

import (
	"log"

	working "github.com/caarlos0/am-i-working"
	"github.com/urfave/cli"
)

func mainAction(c *cli.Context) error {
	events := make(chan bool)
	domain := c.String("domain")
	log.Println("Watching for domain", domain)
	go func() {
		if err := working.Watch(resolv, domain, events); err != nil {
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
