package file

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/YukihiroTaniguchi/pom/domain/model/timeset"
)

// Init ...
func Init(fullPath string, set *timeset.Setting) (err error) {
	f, err := openConfigFileInInit(fullPath)
	defer f.Close()
	if err != nil {
		return err
	}
	fi, err := f.Stat()
	if err != nil {
		return err
	}
	// if already wrote, return
	if fi.Size() > 0 {
		return err
	}
	js, err := json.Marshal(set)
	if err != nil {
		return err
	}
	_, err = f.Write(js)
	return err
}

// Get ...
func Get(fullPath string) (set *timeset.Setting, err error) {
	f, err := openConfigFile(fullPath)
	defer f.Close()
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(f)
	json.Unmarshal(b, &set)
	return
}

// Update ...
func Update(fullPath string, set *timeset.Setting) (err error) {
	f, err := openConfigFile(fullPath)
	defer f.Close()
	if err != nil {
		return err
	}
	js, err := json.Marshal(set)
	if err != nil {
		return err
	}
	_, err = f.Write(js)
	return err
}

func sepPath(p string) (d string, f string) {
	fs := strings.Split(p, "/")
	d = strings.Join(fs[:len(fs)-1], "/")
	f = fs[len(fs)-1]
	return d, f
}

func openConfigFileInInit(p string) (f *os.File, err error) {
	dn, fn := sepPath(p)
	if err = os.Chdir(dn); err != nil {
		os.MkdirAll(dn, 0777)
		os.Chdir(dn)
	}
	f, err = os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0755)
	return f, err
}

func openConfigFile(p string) (f *os.File, err error) {
	dn, fn := sepPath(p)
	if err = os.Chdir(dn); err != nil {
		return
	}
	f, err = os.OpenFile(fn, os.O_RDWR, 0755)
	return
}
