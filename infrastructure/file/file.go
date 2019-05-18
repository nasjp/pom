package file

import (
	"io/ioutil"
	"os"
	"strings"
)

func sepPath(p string) (d string, f string) {
	fs := strings.Split(p, "/")
	d = strings.Join(fs[:len(fs)-1], "/")
	f = fs[len(fs)-1]
	return
}

// Create ...
func Create(json []byte, path string) (err error) {
	dirName, fileName := sepPath(path)
	if err = os.Chdir(dirName); err != nil {
		err = os.MkdirAll(dirName, 0777)
		os.Chdir(dirName)
	}
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	_, err = f.Write(json)
	return
}

// Get ...
func Get(path string) (json []byte, err error) {
	dirName, fileName := sepPath(path)
	if err = os.Chdir(dirName); err != nil {
		return
	}
	f, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer f.Close()
	if err != nil {
		return
	}
	json, err = ioutil.ReadAll(f)
	return
}

// Update ...
func Update(json []byte, path string) (err error) {
	dirName, fileName := sepPath(path)
	if err = os.Chdir(dirName); err != nil {
		return
	}
	f, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer f.Close()
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	if err = f.Truncate(0); err != nil {
		return
	}
	_, err = f.Write(json)
	return
}

// Delete ...
func Delete(path string) (err error) {
	dirName, fileName := sepPath(path)
	if err = os.Chdir(dirName); err != nil {
		return
	}
	err = os.Remove(fileName)
	return
}
