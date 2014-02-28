package main

import "bufio"
import "fmt"
import "os"

func main() {
	stdin := bufio.NewReader(os.Stdin)                   // (1)
	if line, err := stdin.ReadString('\n'); err == nil { // (2)
		fmt.Print(line)
	} else { // (3)
		panic(err) // (4)
	}
	os.Exit(0)
}
