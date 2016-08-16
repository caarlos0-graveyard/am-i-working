# am-i-working [![CircleCI](https://circleci.com/gh/caarlos0/am-i-working.svg?style=svg)](https://circleci.com/gh/caarlos0/am-i-working)

Logs your working activity based on /etc/resolv.conf domain.

The idea is pretty simple: when I'm connected in the company network or
VPN, a line with `domain mycompany.com` appears in my `/etc/resolv.conf`.

What I want here is to log those changes so I can easily get my extra working
hours later (and automate the sending of the report too).

To run it, you can simply `./am-i-working -d mycompany > work.log`, or
create a service in the OS level to keep it running forever.

### OSX Service

Create a file like this:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
	<dict>
		<key>Label</key>
		<string>am-i-working</string>
		<key>ProgramArguments</key>
		<array>
			<string>/path/to/am-i-working</string>
			<string>--domain</string>
			<string>mycompany</string>
		</array>
		<key>RunAtLoad</key>
		<true/>
		<key>StandardOutPath</key>
		<string>/path/to/am-i-working-out.log</string>
		<key>StandardErrorPath</key>
		<string>/path/to/am-i-working-err.log</string>
	</dict>
</plist>
```

Changing the path to the `am-i-working` binary, domain argument and where to
save the logs (usually `/Users/USER/Library/Logs/`).

Save the file in `~/Library/LaunchAgents/am-i-working.plist`.

Then run:

```console
launchctl load ~/Library/LaunchAgents/am-i-working.plist
```

And you can check that the `am-i-working` process will be running on background.
