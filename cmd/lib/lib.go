package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var LogPath string

var logDirPath = "../logs"
var logArchPath = "../logs/arch"

func CreatePathName() string {
	currentTime, err := time.Now().UTC().MarshalText()
	LogPath = "../logs/log_" + strings.Replace(string(currentTime), ":", "", -1) + ".log"
	if err != nil {
		log.Fatalf("Failed to get UTC() time somehow... %v", err)
	}
	return LogPath
}

func PleaseCheckLogError() string {
	return fmt.Sprintf("Error! Please, check log file:\n %s", LogPath)
}

func MoveFilesToArchDir() error {
	//Moves old files to logs/arch directory

	files, err := ioutil.ReadDir(logDirPath)
	if err != nil {
		return err
	}

	// Move each file to the 'Arch' subdirectory
	for _, file := range files {
		if !file.IsDir() {
			oldPath := filepath.Join(logDirPath, file.Name())
			newPath := filepath.Join(logArchPath, file.Name())
			err = os.Rename(oldPath, newPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
