package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runc(cmd string, args []string) string {
	var (
		cmdOut []byte
		cmdErr error
		desc   string
	)
	if cmdOut, cmdErr = exec.Command(cmd, args...).Output(); cmdErr != nil {
		fmt.Fprintln(os.Stderr, cmdErr)
		os.Exit(1) // THINK better things
		fmt.Println(cmd)
	}
	desc = string(cmdOut)
	return desc
}

func main() {
	var (
		out  [3]string
		user string
		host string
	)

	// Collects system-details
	user = os.Getenv("USER")
	host = os.Getenv("HOSTNAME") // Fails try to Run `hostname` instead
	println("Details from " + user + "@" + host)
	// Git log(n)
	n := "1" // TODO
	out[0] = runc("git", []string{"log", "HEAD", "-" + n, "--pretty=short"})
	out[1] = runc("go", []string{"version"})
	out[2] = runc("m-apiserver", []string{"version"})

}
