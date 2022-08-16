package main

type systemctl struct {
	PowerRequester
	application string
}

func newSystemCtl() PowerRequester {
	s := systemctl{application: "/usr/bin/systemctl"}
	return s
}

func (s systemctl) getApplication() string {
	return s.application
}

func (s systemctl) suspend() {
	runCommand(s.application, "suspend")
}

func (s systemctl) hibernate() {
	runCommand(s.application, "hibernate")
}

func (s systemctl) poweroff() {
	runCommand("poweroff")
}

func (s systemctl) reboot() {
	runCommand("reboot")
}
