package main

import "os/exec"

func main() {
	cmd := exec.Command("/usr/bin/iostat", "-xt")
	cmd.Run()
}
