package main

import (
	"flag"
	"log"
	"time"
)

var timeFormat string = "02 Jan 06 15:04 MST"

func parseUserInput(timeValue string) time.Time {
	parsed, err := time.Parse(timeFormat, timeValue)
	log.Println(parsed)
	if err != nil || time.Now().After(parsed) {
		log.Fatal("Error occurred: ", err)
	}
	return parsed
}

func calculateSleepsRemaining(toDate time.Time) float64 {
	return time.Until(toDate).Hours() / 24
}

func main() {
	// just read the date for now and return the date as a time Object
	timeEntry := flag.String("timeValue", "", "The time value I am supposed to process")
	flag.Parse()
	result := parseUserInput(*timeEntry)
	log.Println(result)
	log.Printf("Times to sleep remaining to target date: %d", int(calculateSleepsRemaining(result)))
}
