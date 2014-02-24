package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
)

var subs = map[string]string{}

func substitute(line string) string {
	for orig, replace := range subs {
		re := regexp.MustCompile(orig + "(\\s+)")
		line = re.ReplaceAllString(line, replace+strings.Repeat(" ", 16-len(orig)))
	}
	return line
}

func findDev(rdev uint64) string {
	devdir, err := os.Open("/dev")
	if err != nil {
		panic(err)
	}
	devs, err := devdir.Readdir(0)
	if err != nil {
		panic(err)
	}
	for _, devinfo := range devs {
		if !devinfo.IsDir() {
			if devinfo.Sys().(*syscall.Stat_t).Rdev == rdev {
				return devinfo.Name()
			}
		}
	}
	return ""
}

func generateSubs() {
	file, err := os.Open("/dev/oracleasm/disks")
	if err != nil {
		panic(err)
	}
	disks, err := file.Readdir(0)
	if err != nil {
		panic(err)
	}
	for _, dinfo := range disks {
		rdev := dinfo.Sys().(*syscall.Stat_t).Rdev
		subs[findDev(rdev)] = dinfo.Name()
	}
}

func cmdExec() {
	cmd := exec.Command("iostat", "2")
	stdout, err := cmd.StdoutPipe()
	out := bufio.NewReader(stdout)
	if err != nil {
		os.Exit(-1)
	}
	if err := cmd.Start(); err != nil {
		os.Exit(-1)
	}

	for {
		if buf, err := out.ReadString('\n'); err != nil {
			os.Exit(-1)
		} else {
			fmt.Println(substitute(buf[:len(buf)-1]))
		}
	}
}

func main() {
	generateSubs()
	cmdExec()
}
