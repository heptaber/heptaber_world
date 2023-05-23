package helper

import (
	"log"
	"time"
)

func GetUTCCurrentTimeRFC3339() time.Time {
	currentTime, err := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		log.Fatal("error while parsing current time")
	}
	return currentTime
}
