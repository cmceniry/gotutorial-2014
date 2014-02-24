package mapping

import "time"

func UpdateMappingsRegularly() {
	if err := GenerateSubs(); err != nil {
		panic(err)
	}
	time.Sleep(5*time.Second)
}
