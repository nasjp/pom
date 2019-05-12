package file

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/YukihiroTaniguchi/pom/domain/model/timeset"
)

// InitConfigFile ...
func InitConfigFile(fullPath string, set *timeset.Setting) error {
	if err := cd(fullPath); err != nil {
		return err
	}
	dn, fn := sepPath(fullPath)
	if err := os.Chdir(dn); err != nil {
		os.MkdirAll(dn, 0777)
		os.Chdir(dn)
	}
	f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return err
	}
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

func cd(dir string) error {
	prev, err := filepath.Abs(".")
	if err != nil {
		return err
	}
	defer os.Chdir(prev)
	os.Chdir(dir)
	return err
}

func sepPath(p string) (d string, f string) {
	fs := strings.Split(p, "/")
	d = strings.Join(fs[:len(fs)-1], "/")
	f = fs[len(fs)-1]
	return d, f
}
