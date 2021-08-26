package utils

import (
	"io/ioutil"
)

func GetHostname() (loads string, err error) {
	dat, err := ioutil.ReadFile("/proc/sys/kernel/hostname")
	if err != nil {
		return "", err
	}
	hostname := string(dat)
	return hostname, nil
}
