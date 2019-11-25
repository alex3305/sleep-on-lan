# Sleep on LAN utility

This utility provides a very simple REST service to make a computer go to sleep (or utilize other power states).

## Usage

Run the application as root / Administrator.

## Installation

1. Copy the application (and optionally the unit file) to the target system
2. `sudo cp sleep-on-lan-linux-amd64 /usr/bin/sleep-on-lan`
3. `sudo chmod a+x /usr/bin/sleep-on-lan`
4. Go to `<IP TARGET>:53305` to verify if the application is running

### Installing the unit file

1. `sudo cp sleep-on-lan.service /etc/systemd/system/`
2. `sudo systemctl daemon-reload`
3. `sudo systemctl start sleep-on-lan.service`
4. `sudo systemctl enable sleep-on-lan.service`

## Building

Navigate to the `scripts` directory and run `build.sh` to cross platform build.

## Dependencies

* [httprouter](github.com/julienschmidt/httprouter)
