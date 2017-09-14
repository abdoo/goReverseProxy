package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
)

type Config struct {
	Port      string `json:"port"`
	AppiotURL string `json:"appiot_url"`
}

func GetConfig() *Config{
	file, e := ioutil.ReadFile("ericsson/iot/resources/config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))

	var con Config
	json.Unmarshal(file, &con)
	return &con
}


func main() {
	// Getting Configuration from json file
	config := GetConfig()

	port := config.Port

	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "https",
		Host:   config.AppiotURL,

	})

	director := proxy.Director
	proxy.Director = func(req *http.Request) {
		director(req)
		req.Host = req.URL.Host
	}


	fmt.Printf("starting server\n")

	log.Fatal(http.ListenAndServe(":"+port, proxy))
}

