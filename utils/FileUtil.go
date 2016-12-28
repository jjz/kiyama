package utils

import (
	"os"
	"io/ioutil"
	"path/filepath"
	"os/exec"
	"strings"
	"fmt"
	"crypto/md5"
	"io"
)

func ReadFile(path string) (str string, err error) {
	fi, err := os.Open(path)
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	str = string(fd)
	return

}
func ReadFileByte(path string) ([]byte, error) {
	fi, err := os.Open(path)
	if err != nil {
		fmt.Println("err", err)
	}
	defer fi.Close()
	return ioutil.ReadAll(fi)

}

func GetAppRoot() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	p, err := filepath.Abs(file)
	if err != nil {
		return ""
	}
	return filepath.Dir(p)

}

func MergePath(args ...string) string {
	approot := GetAppRoot()
	for i, e := range args {
		if e != "" {
			return filepath.Join(approot, filepath.Clean(strings.Join(args[i:], string(filepath.Separator))))

		}
	}
	return approot
}
func GetMd5FromFile(filePath string) (string) {
	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	md5h := md5.New()
	io.Copy(md5h, file)
	return fmt.Sprintf("%x", md5h.Sum([]byte("")))

}