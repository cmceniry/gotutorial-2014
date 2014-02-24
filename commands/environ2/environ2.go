package main

import "fmt"
import "syscall"

func main() {
	fmt.Println(syscall.Environ()[0])       // (1)
	for _, val := range syscall.Environ() { // (2)
		fmt.Println(val)
	}
	fmt.Println(syscall.Environ()[10000]) // (3)
}
