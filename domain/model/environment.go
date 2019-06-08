package model

import (
	"fmt"
	"os"
)

var (
	gopath     = "GOPATH"
	appDir     = "/src/github.com/NasSilverBullet/pom"
	configFile = "/config/timerset.json"
	grobalFile = "/config/grobal.json"
)

// GetGOPATH ...
func GetGOPATH() (ep string, err error) {
	ep = os.Getenv(gopath)
	if ep == "" {
		err = fmt.Errorf("\"%s\" is not defined, please define", gopath)
	}
	return
}

// GetAppDir ...
func GetAppDir() string {
	return appDir
}

// GetConfigFile ...
func GetConfigFile() string {
	return configFile
}

// GetGrobalFile ...
func GetGrobalFile() string {
	return grobalFile
}
