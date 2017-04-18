package commands

import (
	"github.com/caarlos0/am-i-working/service"
	"github.com/urfave/cli"
)

// CreateService creates the service on the OS
var CreateService = cli.Command{
	Name:  "create",
	Usage: "creates the service on the OS",
	Flags: flags,
	Action: func(c *cli.Context) error {
		var resolv = c.String("file")
		var domain = c.String("domain")
		if domain == "" {
			return cli.NewExitError("missing domain name", 1)
		}
		return service.Create(resolv, domain)
	},
}

// StartService starts the service on the OS
var StartService = cli.Command{
	Name:  "start",
	Usage: "starts the service on the OS",
	Action: func(c *cli.Context) error {
		return service.Start()
	},
}

// StopService stops the service on the OS
var StopService = cli.Command{
	Name:  "stop",
	Usage: "stops the service on the OS",
	Action: func(c *cli.Context) error {
		return service.Stop()
	},
}

// RestartService stops and starts the service
var RestartService = cli.Command{
	Name:  "restart",
	Usage: "restarts the service",
	Action: func(c *cli.Context) error {
		if err := service.Stop(); err != nil {
			return err
		}
		return service.Start()
	},
}

// DeleteService stops and delete the service
var DeleteService = cli.Command{
	Name:  "delete",
	Usage: "stops and delete the service",
	Action: func(c *cli.Context) error {
		_ = service.Stop()
		return service.Delete()
	},
}
