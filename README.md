# am-i-working

Logs your working activity based on /etc/resolv.conf domain.

The idea is pretty simple: when I'm connected in the company network or
VPN, a line with `domain mycompany.com` appears in my `/etc/resolv.conf`.

What I want here is to log those changes so I can easily get my extra working
hours later (and automate the sending of the report too).

To run it, you can simply `./am-i-working -d mycompany > work.log`, or 
create a service in the OS level to keep it running forever.
