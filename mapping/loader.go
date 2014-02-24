package mapping // (1)

import "bufio"
import "fmt"
import "os"

func Read(name string) (map[string]string, error) { // (2)
	file, err := os.Open(name) // (3)
	if err != nil {
		return nil, err // (4)
	}
	defer file.Close() 
	ret := make(map[string]string)
	bio := bufio.NewReader(file)
	for {
		var src, dst string
		if line, err := bio.ReadString('\n'); err == nil {
			if _, err := fmt.Sscanf(line, "%s %s", &src, &dst); err == nil {
				ret[src] = dst
			}
		} else {
			break
		}
	}
	return ret, nil
}
