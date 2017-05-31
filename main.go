package main

import (
	"flag"
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
		fmt.Println(cmd)
		os.Exit(1) // THINK better things

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

	i_commits := flag.String("commits", "1", "an int")
	i_logs := flag.String("logfile", "", "a string")

	flag.Parse()

	// Collects system-details
	user = os.Getenv("USER")
	host = runc("hostname", []string{"-s"})
	fmt.Println(user, host)
	out[0] = runc("git", []string{"log", "HEAD", "-" + *i_commits, "--pretty=short"})
	out[1] = runc("go", []string{"version"})
	out[2] = runc("m-apiserver", []string{"version"})

	// Collect repo details
	// FIXME What if a project has multiple remotes

	if *i_logs != "" {
		if _, err := os.Stat(*i_logs); os.IsNotExist(err) {
			fmt.Fprintln(os.Stderr, *i_logs+" not found. Issue will be created without the logs")
		}
	}

}
