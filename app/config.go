package app

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var Config config

type config struct {
	Server struct {
		RunMode      string
		HttpPort     int64
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}
	DB struct {
		Host     string
		Port     string
		DBType   string
		User     string
		Password string
		Name   string
	}
	Version struct {
		Release string
	}
	App struct {
		LogSavePath     string
		LogSaveName     string
		LogFileExt      string
		RuntimeRootPath string
		LogTimeFormat   string
	}
}

var (
	cfg            *ini.File
)

func InitConfig() {
	var err error
	cfg, err = ini.Load("conf/conf.ini")
	if err != nil {
		log.Fatalf("settting.Setup, fail to load conf file:%v", err)
	}
	sectionMapTo("server", &Config.Server)
	sectionMapTo("app", &Config.App)
	sectionMapTo("database", &Config.DB)
	sectionMapTo("version", &Config.Version)
	Config.Server.ReadTimeout = Config.Server.ReadTimeout*time.Second
	Config.Server.WriteTimeout = Config.Server.WriteTimeout*time.Second
}

func sectionMapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("cfg.sectionMapTo err: %v", err)
	}
}
