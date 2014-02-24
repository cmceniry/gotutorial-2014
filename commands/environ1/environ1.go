package main

import "fmt"
import "syscall"

func main() {
	if val, found := syscall.Getenv("GOPATH"); found {
		fmt.Println(val)
	}
}
