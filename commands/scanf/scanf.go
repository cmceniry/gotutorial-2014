package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

var subs = map[string]string{
	"dm-2": "ASM001",
	"dm-3": "ASM002",
	"dm-4": "ASM003",
	"dm-5": "ASM004",
	"dm-6": "ASM005",
	"dm-7": "ASM006",
	"dm-8": "ASM007",
	"dm-9": "ASM008",
}

func main() {
	cmd := exec.Command("iostat", "-xt", "2")
	stdout, err := cmd.StdoutPipe()
	out := bufio.NewReader(stdout)
	if err != nil {
		os.Exit(-1)
	}
	if err := cmd.Start(); err != nil {
		os.Exit(-1)
	}

	var ts string

	var mm, dd, yy, h, m, s int32
	var ap string
	tsin := "%d/%d/%d %d:%d:%d %s" // (1)
	tsout := "%02d/%02d/%d:%02d:%02d:%02d%s"

	var device string
	var rs, ws, srs, sws, ign float32
	devin := "%s %f %f %f %f %f %f %f %f %f %f %f"
	devout := "%s|%s|%.2f|%.2f|%.2f|%.2f\n"

	for {
		if buf, err := out.ReadString('\n'); err != nil {
			os.Exit(-1)
		} else {
			if _, err := fmt.Sscanf(buf, tsin, &mm, &dd, &yy, &h, &m, &s, &ap); err == nil { // (2)
				ts = fmt.Sprintf(tsout, mm, dd, yy, h, m, s, ap) // (3)
				continue                                         // (4)
			}
			if _, err := fmt.Sscanf(buf, devin, &device, &ign, &rs, &ws, &srs, &sws, &ign, &ign, &ign, &ign, &ign, &ign); err == nil { // (5)
				if _, ok := subs[device]; ok {
					fmt.Printf(devout, ts, subs[device], rs, ws, srs, sws)
				} else {
					fmt.Printf(devout, ts, device, rs, ws, srs, sws)
				}
			}
		}
	}
}
