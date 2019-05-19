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

func create(json []byte, path string) (err error) {
	dirName, fileName := sepPath(path)
	if err = os.Chdir(dirName); err != nil {
		if err = os.MkdirAll(dirName, 0777); err != nil {
			return
		}
		if err = os.Chdir(dirName); err != nil {
			return
		}
	}
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	defer func() {
		err = f.Close()
	}()
	if err != nil {
		return
	}
	_, err = f.Write(json)
	return
}

func get(path string) (json []byte, err error) {
	dirName, fileName := sepPath(path)
	if err = os.Chdir(dirName); err != nil {
		return
	}
	f, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer func() {
		err = f.Close()
	}()
	json, err = ioutil.ReadAll(f)
	return
}

func update(json []byte, path string) (err error) {
	dirName, fileName := sepPath(path)
	if err = os.Chdir(dirName); err != nil {
		return
	}
	f, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer func() {
		err = f.Close()
	}()
	if err = f.Truncate(0); err != nil {
		return
	}
	_, err = f.Write(json)
	return
}

func remove(path string) (err error) {
	dirName, fileName := sepPath(path)
	if err = os.Chdir(dirName); err != nil {
		return
	}
	err = os.Remove(fileName)
	return
}
