# am-i-working [![Build Status](https://travis-ci.org/caarlos0/am-i-working.svg?branch=master)](https://travis-ci.org/caarlos0/am-i-working) [![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)

Logs your working activity based on /etc/resolv.conf domain.

The idea is pretty simple: when I'm connected in the company network or
VPN, a line with `domain mycompany.com` appears in my `/etc/resolv.conf`.

What I want here is to log those changes so I can easily get my extra working
hours later (and automate the sending of the report too).

To run it, you can simply `./am-i-working watch --domain mycompany > work.log`,
or create a service in the OS level to keep it running forever.

## Install

```console
brew install caarlos0/tap/am-i-working
```

### macOS Service

To set it up as a macOS service, just tun:

```console
$ am-i-working create --domain mycompany
$ am-i-working start
```

There are also the `stop` and `restart` commands.
