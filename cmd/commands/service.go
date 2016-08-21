package commands

import "github.com/urfave/cli"

// CreateService creates the service on the OS
var CreateService = cli.Command{
	Name:  "create",
	Usage: "creates the service on the OS",
	Flags: flags,
	Action: func(c *cli.Context) error {
		return create(c)
	},
}

// StartService starts the service on the OS
var StartService = cli.Command{
	Name:  "start",
	Usage: "starts the service on the OS",
	Action: func(c *cli.Context) error {
		return start(c)
	},
}

// StopService stops the service on the OS
var StopService = cli.Command{
	Name:  "stop",
	Usage: "stops the service on the OS",
	Action: func(c *cli.Context) error {
		return stop(c)
	},
}

// RestartService stops and starts the service
var RestartService = cli.Command{
	Name:  "restart",
	Usage: "restarts the service",
	Action: func(c *cli.Context) error {
		if err := stop(c); err != nil {
			return err
		}
		return start(c)
	},
}
