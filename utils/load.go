package utils

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseLoad(l string) (datas []byte, err error) {
	load := strings.Fields(string(l))
	m := make(map[string]float32)
	load1m, _ := strconv.ParseFloat(load[0], 32)
	load5m, _ := strconv.ParseFloat(load[1], 32)
	load15m, _ := strconv.ParseFloat(load[2], 32)

	m["1m"] = float32(load1m)
	m["5m"] = float32(load5m)
	m["15m"] = float32(load15m)
	loadJson, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return loadJson, err
}

func GetLoad() (loads string, err error) {
	dat, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		return "", err
	}
	// {"1m": 0.57, "5m":0.80, "15m":0.77}
	loadJson, err := parseLoad(string(dat))
	if err != nil {
		return "", err
	}
	load := string(loadJson)
	return load, nil
}
