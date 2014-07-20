package main

import "bufio"
import "flag"
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
	var dontprocess bool
	f := flag.NewFlagSet("myname", flag.ExitOnError) // (1)
	f.BoolVar(&dontprocess, "dontprocess", false,
		"Disables substitution of device names") // (2)
	f.Parse(os.Args[1:]) // (3)

	bio := bufio.NewReader(os.Stdin)
	done := false
	for !done {
		if line, err := bio.ReadString('\n'); err == nil {
			if !dontprocess {
				fmt.Print(substitute(line))
			} else {
				fmt.Print(line)
			}
		} else {
			done = true
		}
	}
}
