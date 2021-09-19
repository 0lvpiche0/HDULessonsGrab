package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Base    map[string]string   `yaml:"base"`
	Headers map[string][]string `yaml:"headers"`
	Body    string              `yaml:"body"`
}

var config Config

func cfgInit() {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		panic(err)
	}
}

type HttpCfg struct {
	Url     string
	Method  string
	Body    io.Reader
	Headers map[string][]string
}

func NewHttpCfg(config Config, method string) HttpCfg {
	return HttpCfg{
		Url:     config.Base["baseurl"] + "/jwglxt/xsxk/zzxkyzbjk_xkBcZyZzxkYzb.html?gnmkdm=N253512&su=" + config.Base["su"],
		Method:  method,
		Body:    strings.NewReader(config.Body),
		Headers: config.Headers,
	}
}

type ResponseData struct {
	Flag string
	Msg  string
}

func main() {
	cfgInit()
	httpCfg := NewHttpCfg(config, "POST")
	req, err := http.NewRequest(httpCfg.Method, httpCfg.Url, httpCfg.Body)
	if err != nil {
		panic(err)
	}
	req.Header = httpCfg.Headers
	client := &http.Client{}
	for {
		// fmt.Println(req)
		response, err := client.Do(req)
		if err != nil {
			log.Println(err)
			return
		}
		data, err := io.ReadAll(response.Body)
		// fmt.Println(string(data))
		if err != nil {
			log.Println(err)
			return
		}
		var responseData ResponseData
		if err := json.Unmarshal(data, &responseData); err != nil {
			log.Println(err)
			return
		}
		if responseData.Flag != "-1" {
			log.Println("flag != -1\n", responseData)
			return
		}
		time.Sleep(30 * time.Second)
	}
}
