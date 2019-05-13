package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/YukihiroTaniguchi/pom/domain/model/timeset"
)

const (
	// GOPATH ...
	GOPATH = "GOPATH"
	// APPDIR ...
	APPDIR = "/src/github.com/YukihiroTaniguchi/pom"
	// CONFIGFILE ...
	CONFIGFILE = "/config/pom.json"
)

var (
	defaultConfig = timeset.Setting{
		Work:       25,
		ShortBreak: 10,
		LongBreak:  20,
		Set:        10,
	}
)

// Init ...
func Init() (err error) {
	f, err := openConfigFileInInit()
	defer f.Close()
	if err != nil {
		return
	}
	fi, err := f.Stat()
	if err != nil {
		return
	}
	// if already wrote, return
	if fi.Size() > 0 {
		return
	}
	js, err := json.Marshal(defaultConfig)
	if err != nil {
		return
	}
	_, err = f.Write(js)
	return
}

// Get ...
func Get() (set *timeset.Setting, err error) {
	f, err := openConfigFile()
	defer f.Close()
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(f)
	json.Unmarshal(b, &set)
	return
}

// Update ...
func Update(set *timeset.Setting) (err error) {
	f, err := openConfigFile()
	defer f.Close()
	if err != nil {
		return
	}
	js, err := json.Marshal(set)
	if err != nil {
		return
	}
	_, err = f.Write(js)
	return
}

func getEnvVar() (ep string, err error) {
	ep = os.Getenv(GOPATH)
	if ep == "" {
		err = fmt.Errorf("\"%s\" is not defined, please define", GOPATH)
	}
	return
}

func sepPath(p string) (d string, f string) {
	fs := strings.Split(p, "/")
	d = strings.Join(fs[:len(fs)-1], "/")
	f = fs[len(fs)-1]
	return
}

func openConfigFileInInit() (f *os.File, err error) {
	ep, err := getEnvVar()
	if err != nil {
		return
	}
	p := ep + APPDIR + CONFIGFILE
	dn, fn := sepPath(p)
	if err = os.Chdir(dn); err != nil {
		os.MkdirAll(dn, 0777)
		os.Chdir(dn)
	}
	f, err = os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0755)
	return
}

func openConfigFile() (f *os.File, err error) {
	ep, err := getEnvVar()
	if err != nil {
		return
	}
	p := ep + APPDIR + CONFIGFILE
	dn, fn := sepPath(p)
	if err = os.Chdir(dn); err != nil {
		return
	}
	f, err = os.OpenFile(fn, os.O_RDWR, 0755)
	return
}
