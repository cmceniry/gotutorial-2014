package main

import "os/exec"

func main() {
	args := []string{"-xt", "/dev/loop1", "/dev/loop2", "/dev/loop3", "/dev/loop4"}
	cmd := exec.Command("/usr/bin/iostat", args...)
	cmd.Run()
}
