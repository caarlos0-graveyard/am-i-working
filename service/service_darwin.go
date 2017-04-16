// Package service contains am-i-working OS service implementations
package service

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/kardianos/osext"
)

const xmlTemplate = `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
	<dict>
		<key>Label</key>
		<string>am-i-working</string>
		<key>ProgramArguments</key>
		<array>
			<string>{{ .Binary }}</string>
			<string>watch</string>
			<string>--domain</string>
			<string>{{ .Domain }}</string>
			<string>--file</string>
			<string>{{ .Resolv }}</string>
		</array>
		<key>RunAtLoad</key>
		<true/>
		<key>StandardOutPath</key>
		<string>{{ .Home }}/Library/Logs/am-i-working-out.log</string>
		<key>StandardErrorPath</key>
		<string>{{ .Home }}/Library/Logs/am-i-working-out.log</string>
	</dict>
</plist>
`

var plist = filepath.Join(os.Getenv("HOME"), "Library", "LaunchAgents", "am-i-working.plist")

// Create the service
func Create(resolvFile, domainName string) error {
	var out bytes.Buffer
	t, err := template.New("service").Parse(xmlTemplate)
	if err != nil {
		return err
	}
	executable, err := osext.Executable()
	if err != nil {
		return err
	}

	err = t.Execute(&out, struct {
		Binary, Domain, Resolv, Home string
	}{
		Binary: executable,
		Domain: domainName,
		Resolv: resolvFile,
		Home:   os.Getenv("HOME"),
	})
	if err != nil {
		return err
	}
	fmt.Println("writing", plist)
	return ioutil.WriteFile(plist, out.Bytes(), 0644)
}

// Delete the service
func Delete() error {
	fmt.Println("removing", plist)
	return os.Remove(plist)
}

// Start the service
func Start() error {
	return exec.Command("launchctl", "load", plist).Run()
}

// Stop the service
func Stop() error {
	return exec.Command("launchctl", "unload", plist).Run()
}
