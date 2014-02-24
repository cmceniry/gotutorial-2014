package main

import (
    "bufio"
    "os"
    "fmt"
)

func main() {
    bio := bufio.NewReader(os.Stdin)
    for { // (1)
        if line, err := bio.ReadString('\n'); err == nil {
            fmt.Print(line)
        } else {
            break // (2)
        }
    }

}
