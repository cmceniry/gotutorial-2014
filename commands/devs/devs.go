package main

import "github.com/cmceniry/gotutorial/mapping"
import "bufio"
import "fmt"
import "os/exec"
import "regexp"
import "strings"

func substitute(line string) string {
	for orig, replace := range mapping.Subs {
		re := regexp.MustCompile(orig + "(\\s+)")
		line = re.ReplaceAllString(line, replace+strings.Repeat(" ", 16-len(orig)))
	}
	return line
}

func cmdExec() {
	cmd := exec.Command("iostat", "-xt", "5")
	stdout, err := cmd.StdoutPipe()
	out := bufio.NewReader(stdout)
	if err != nil {
		panic(err)
	}
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	for {
		if buf, err := out.ReadString('\n'); err != nil {
			panic(err)
		} else {
			fmt.Print(substitute(buf))
		}
	}
}

func main() {
	mapping.GenerateSubs()
	cmdExec()
}
