package file

import (
	"io/ioutil"
	"os"
)

/**
* @Author: super
* @Date: 2020-09-23 07:56
* @Description:
**/

func ReadFile(path string) (string, error) {
	fi, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fi.Close()

	fd, err := ioutil.ReadAll(fi)
	return string(fd), err
}