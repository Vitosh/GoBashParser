package lib

import (
	"log"
	"strings"
	"time"
)

var LogPath string

func CreatePath() string {
	currentTime, err := time.Now().UTC().MarshalText()
	LogPath = "../logs/log_" + strings.Replace(string(currentTime), ":", "", -1) + ".log"

	if err != nil {
		log.Fatalf("Failed to get UTC() time somehow... %v", err)
	}
	return LogPath
}
