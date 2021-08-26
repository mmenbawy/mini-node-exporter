package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mmenbawy/mini-node-exporter/utils"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/info/hostname" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	hostname, _ := utils.GetHostname()
	fmt.Fprintf(w, hostname)
}

func uptimeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/info/uptime" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	uptime, _ := utils.GetUptime()
	fmt.Fprintf(w, uptime)
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/info/load" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	load, _ := utils.GetLoad()
	fmt.Fprintf(w, load)
}

func main() {
	utils.SetUptime()
	utils.SetLoad()

	http.HandleFunc("/info/hostname", hostnameHandler)
	http.HandleFunc("/info/uptime", uptimeHandler)
	http.HandleFunc("/info/load", loadHandler)
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":23333", nil); err != nil {
		log.Fatal(err)
	}
}
