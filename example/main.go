package main

import (
	".."
	"log"
)


func main() {
	loadavg, err := loadavg.Parse()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("load averages: %.2f %.2f %.2f", loadavg.LoadAverage1, loadavg.LoadAverage5, loadavg.LoadAverage10)
}