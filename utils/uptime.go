package utils

import (
	"io/ioutil"
	"strings"
)

func GetUptime() (loads string, err error) {
	dat, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		return "", err
	}
	uptime := strings.Fields(string(dat))
	return uptime[0], nil
}
