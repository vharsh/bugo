package main

import (
	"os/exec"
)

func runc(cmd, args string) {
	cmd := exec.Command(cmd, args)
}

func main() {
	runc("git", "log HEAD -1 --pretty=short") // TODO Modify it to take user# of commits
}
