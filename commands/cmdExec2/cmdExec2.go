package main

import "fmt"
import "os/exec"
import "time"

func main() {
	buf := make([]byte, 2048)
	cmd := exec.Command("/usr/bin/iostat", "-xt")
	stdoutpipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	if err := cmd.Start(); err != nil { // (1)
		panic(err)
	}
	time.Sleep(1 * time.Second)    // (2)
	n, err := stdoutpipe.Read(buf) // (3)
	if err != nil {
		panic(err)
	}
	cmd.Wait() // (4)

	fmt.Print(string(buf[:n])) // (5)
}
