package main

import "github.com/cmceniry/gotutorial/mapping" // (1)
import "bufio"
import "fmt"
import "os"
import "regexp"

var subs = map[string]string{}

func substitute(before string) string {
	for needle, sub := range subs {
		if re, err := regexp.Compile(needle); err == nil {
			before = re.ReplaceAllString(before, sub)
		}
	}
	return before
}

func main() {
	if cf, err := mapping.Read("/home/golang/mapping"); err != nil { // (2)
		panic(err)
	} else {
		subs = cf
	}

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
