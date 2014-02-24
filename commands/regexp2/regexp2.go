package main

import (
    "bufio"
    "os"
    "regexp"
    "fmt"
)

var subs = map[string]string{
    "dm-2  ": "ASM001",
    "dm-3  ": "ASM002",
    "dm-4  ": "ASM003",
    "dm-5  ": "ASM004",
    "dm-6  ": "ASM005",
    "dm-7  ": "ASM006",
    "dm-8  ": "ASM007",
    "dm-9  ": "ASM008",
}

func substitute(before string) string {
    for needle, sub := range subs { // (1)
        if re, err := regexp.Compile(needle); err == nil {
            before = re.ReplaceAllString(before, sub) // (2)
        }
    }
    return before
}

func main() {
    bio := bufio.NewReader(os.Stdin)
    done := false
    for !done {
        if line, err := bio.ReadString('\n'); err == nil {
            fmt.Print(substitute(line))
        } else {
            done = true
        }
    }
}

