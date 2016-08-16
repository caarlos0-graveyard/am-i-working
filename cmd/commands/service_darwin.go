package commands

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/kardianos/osext"
	"github.com/urfave/cli"
)

const xml = `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
	<dict>
		<key>Label</key>
		<string>am-i-working</string>
		<key>ProgramArguments</key>
		<array>
			<string>BINARY</string>
			<string>watch</string>
			<string>--domain</string>
			<string>DOMAIN</string>
			<string>--file</string>
			<string>FILE</string>
		</array>
		<key>RunAtLoad</key>
		<true/>
		<key>StandardOutPath</key>
		<string>HOME/Library/Logs/am-i-working-out.log</string>
		<key>StandardErrorPath</key>
		<string>HOME/Library/Logs/am-i-working-out.log</string>
	</dict>
</plist>
`

var plist = os.Getenv("HOME") + "/Library/LaunchAgents/am-i-working.plist"

func create(c *cli.Context) error {
	executable, err := osext.Executable()
	if err != nil {
		return err
	}
	svc := strings.Replace(xml, "FILE", c.String("file"), -1)
	svc = strings.Replace(svc, "HOME", os.Getenv("HOME"), -1)
	svc = strings.Replace(svc, "DOMAIN", c.String("domain"), -1)
	svc = strings.Replace(svc, "BINARY", executable, -1)
	return ioutil.WriteFile(plist, []byte(svc), 0644)
}

func start(c *cli.Context) error {
	return exec.Command("launchctl", "load", plist).Run()
}

func stop(c *cli.Context) error {
	return exec.Command("launchctl", "unload", plist).Run()
}
