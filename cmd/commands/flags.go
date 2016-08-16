package commands

import "github.com/urfave/cli"

var flags = []cli.Flag{
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
