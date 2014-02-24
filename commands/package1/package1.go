package main

import (
	"github.com/cmceniry/gotutorial/v1/inCmd"
	"github.com/cmceniry/gotutorial/v1/inStdin"
	"github.com/cmceniry/gotutorial/v1/subRegexp"
	"github.com/cmceniry/gotutorial/v1/subScanf"
	"flag"
)

func main() {
	var dontprocess bool
	var useStdin bool
	var output string
	f := flag.NewFlagSet("mine", flag.ExitOnError)
	f.BoolVar(&dontprocess, "dontprocess", false, "Don't do any processing")
	f.BoolVar(&useStdin, "use-stdin", false, "Use stdin instead of cmd execution")
	f.StringVar(&output, "output", "regexp", "Which output type: regexp or scanf")
	f.Parse(os.Args[1:])

	if !useStdin {
		inCmd.Start()
	}

	var buf string
	for {
		if useStdin {
			buf = inStdin.GetLine()
		} else {
			buf = inCmd.GetLine()
		}
		if dontProcess {
			subRegexp.Print(buf, dontProcess)
		} else {
			switch output {
			case "regexp": subRegexp.Print(buf, dontProcess)
			case "scanf": subScanf.Print(scanf)
		}
	}
}
