# DigitalOcean DDNS client

## Description
This client is useful when you have the domain managed by DigitalOcean and want 
to set the IP address it should point to.
Personally I use it in my Asus RT-AC68U (ARM based) router as a custom dymanic DNS client.

## Compiling
There is a `Makefile` to build the application for ARM (Linux) and your current platform.
Simply run `make`.

## Using
You have to obtain Digital Ocean OAuth token first.

See https://developers.digitalocean.com/documentation/v2 for more info.

Simply run `do-ddns -ip {YOUR IP ADDRESS} -host {YOUR DOMAIN NAME} -token {YOUR DIGITAL OCEAN TOKEN}`