package mapping

import "os"
import "syscall"

var Subs map[string]string

func findDev(rdev uint64) (string, error) {
	devdir, err := os.Open("/dev")
	if err != nil {
		return "", err
	}
	devs, err := devdir.Readdir(0)
	if err != nil {
		return "", err
	}
	for _, devinfo := range devs {
		if !devinfo.IsDir() {
			if devinfo.Sys().(*syscall.Stat_t).Rdev == rdev {
				return devinfo.Name(), nil
			}
		}
	}
	return "", nil
}

func GenerateSubs() error {
	file, err := os.Open("/dev/oracleasm/disks") // (1)
	if err != nil {
		return err
	}
	disks, err := file.Readdir(0) // (2)
	if err != nil {
		return err
	}
	var subs = make(map[string]string)
	for _, dinfo := range disks {
		rdev := dinfo.Sys().(*syscall.Stat_t).Rdev // (3)
		if dname, err := findDev(rdev); err == nil { // (4)
			subs[dname] = dinfo.Name()
		}
	}
	Subs = subs
	return nil
}
