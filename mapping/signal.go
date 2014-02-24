package mapping

import "time"

func SignalUpdater(done chan bool) { // (1)
	var first = true
	for {
		if err := GenerateSubs(); err != nil {
			panic(err)
		}
		if first { // (2)
			done <- true // (3)
			first = false
		}
		time.Sleep(5*time.Second)
	}
}
