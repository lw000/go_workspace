package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/jimlawless/cfg"
)

var (
	Conf *Config

//	confFile string
)

var confFile = *flag.String("config_file", "./cofnfig.conf", " please sset config file path...")

//func init() {
//	flag.StringVar(&confFile, "c", "./cofnfig.conf", " set config file path")
//}

type Config struct {
	HostAddress string
	ApiKey      string
}

func NewConfig() *Config {
	return &Config{
		HostAddress: "127.0.0.1:5008",
		ApiKey:      "",
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
	Conf.HostAddress = getString(confMap, "HostAddress")
	Conf.ApiKey = getString(confMap, "ApiKey")

	log.Println("ApiKey:", Conf.ApiKey)
	log.Println("HostAddress:", Conf.HostAddress)
}

func getInt(cfg map[string]string, key string) int {
	value, ok := cfg[key]
	if !ok {
		log.Fatalf("key:%s non exist", key)
	}

	n, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("key:%s is't interger", key)
	}
	return n
}

func getString(cfg map[string]string, key string) string {
	value, ok := cfg[key]
	if !ok {
		log.Fatalf("key:%s non exist", key)
	}
	return value
}

func getOptString(cfg map[string]string, key string) string {
	value, ok := cfg[key]
	if !ok {
		return ""
	}
	return value
}
