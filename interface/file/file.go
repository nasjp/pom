package file

import (
	"fmt"
	"os"
	"strings"
)

// CreateOrNot ...
func CreateOrNot(filePath string) error {
	dn, fn := sepPath(filePath)
	if err := os.Chdir(dn); err != nil {
		os.MkdirAll(dn, 0777)
		os.Chdir(dn)
	}
	f, err := os.Open(fn)
	if err != nil {
		f, err = os.Create(fn)
	}
	defer f.Close()
	return err
}

func sepPath(p string) (d string, f string) {
	fs := strings.Split(p, "/")
	d = strings.Join(fs[:len(fs)-1], "/")
	f = fs[len(fs)-1]
	fmt.Println(d, f)
	return d, f
}
