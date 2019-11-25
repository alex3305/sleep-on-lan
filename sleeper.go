package main

import (
	"log"
	"os/exec"
)

// PowerRequester implements power state requests from the target system.
type PowerRequester interface {
	hibernate()
	suspend()
	poweroff()
	reboot()
	getApplication() string
}

func runCommand(cmd string, params ...string) {
	c := exec.Command(cmd, params...)
	err := c.Start()
	
	if err != nil {
		log.Println(err)
	}

	// We never wait here for the command to complete. This would be a
	// waste of time, since the computer is going to another power state
	// anyway. 
}
