# am-i-working

[![Release](https://img.shields.io/github/release/caarlos0/am-i-working.svg?style=flat-square)](https://github.com/caarlos0/am-i-working/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Travis](https://img.shields.io/travis/caarlos0/am-i-working.svg?style=flat-square)](https://travis-ci.org/caarlos0/am-i-working)
[![Coverage Status](https://img.shields.io/coveralls/caarlos0/am-i-working/master.svg?style=flat-square)](https://coveralls.io/github/caarlos0/am-i-working?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/caarlos0/am-i-working?style=flat-square)](https://goreportcard.com/report/github.com/caarlos0/am-i-working)
[![Godoc](https://godoc.org/github.com/caarlos0/am-i-working?status.svg&style=flat-square)](http://godoc.org/github.com/caarlos0/am-i-working)
[![SayThanks.io](https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg?style=flat-square)](https://saythanks.io/to/caarlos0)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)

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
