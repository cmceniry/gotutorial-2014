package main

import "bufio"
import "fmt"
import "os"
import "regexp"

var subs = map[string]string{
	"loop1 ": "ASM001",
	"loop2 ": "ASM002",
	"loop3 ": "ASM003",
	"loop4 ": "ASM004",
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
