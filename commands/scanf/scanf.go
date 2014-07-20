package main

import "bufio"
import "fmt"
import "os"
import "os/exec"

var subs = map[string]string{
	"loop1": "ASM001",
	"loop2": "ASM002",
	"loop3": "ASM003",
	"loop4": "ASM004",
}

func main() {
	args := []string{"-xt", "/dev/loop1", "/dev/loop2", "/dev/loop3", "/dev/loop4", "2"}
	cmd := exec.Command("iostat", args...)
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
		if line, err := out.ReadString('\n'); err != nil {
			os.Exit(-1)
		} else {
			if _, err := fmt.Sscanf(line, tsin, &mm, &dd, &yy, &h, &m, &s, &ap); err == nil { // (2)
				ts = fmt.Sprintf(tsout, mm, dd, yy, h, m, s, ap) // (3)
				continue                                         // (4)
			}
			if _, err := fmt.Sscanf(line, devin, &device, &ign, &rs, &ws, &srs, &sws, &ign, &ign, &ign, &ign, &ign, &ign); err == nil { // (5)
				if _, ok := subs[device]; ok {
					fmt.Printf(devout, ts, subs[device], rs, ws, srs, sws)
				} else {
					fmt.Printf(devout, ts, device, rs, ws, srs, sws)
				}
			}
		}
	}
}
