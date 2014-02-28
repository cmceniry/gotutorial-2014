package main


import "bufio"
import "fmt"
import "os"
import "regexp"

func substitute(before string) string { // (1)
	if re, err := regexp.Compile("dm-2  "); err == nil { // (2)
		return re.ReplaceAllString(before, "ASM001") // (3)
	}
	return before
}

func main() {
	bio := bufio.NewReader(os.Stdin)
	done := false
	for !done {
		if line, err := bio.ReadString('\n'); err == nil {
			fmt.Print(substitute(line)) // (4)
		} else {
			done = true
		}
	}
}
