package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/jimlawless/cfg"
)

var (
	Conf     *Config
	confFile string
)

func init() {
	flag.StringVar(&confFile, "c", "./cofnfig.conf", " set config file path")
}

type Config struct {
	// base section
	host_address string
	api_key      string
}

func NewConfig() *Config {
	return &Config{
		host_address: "127.0.0.1:5008",
		api_key:      "",
	}
}

// 初始化全局配置.
func InitConfig() {
	Conf = NewConfig()
	confMap := make(map[string]string)
	err := cfg.Load(confFile, confMap)
	if err != nil {
		log.Fatal(err)
	}
	Conf.host_address = getString(confMap, "host_address")
	Conf.api_key = getString(confMap, "api_key")

	log.Println("api_key:", Conf.api_key)
	log.Println("host address:", Conf.host_address)
}

func getInt(config map[string]string, key string) int {
	value, ok := config[key]
	if !ok {
		log.Fatalf("key:%s non exist", key)
	}
	n, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("key:%s is't interger", key)
	}
	return n
}

func getString(config map[string]string, key string) string {
	value, ok := config[key]
	if !ok {
		log.Fatalf("key:%s non exist", key)
	}
	return value
}

func getOptString(config map[string]string, key string) string {
	value, ok := config[key]
	if !ok {
		return ""
	}
	return value
}
