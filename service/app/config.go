package app

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"log"
	"os"
	"path/filepath"
	"strings"
	"yuanzhoucehui.com/be/service/extensions"
)

type Config struct {
	rest.RestConf
	LOG extensions.LogConf `json:"zlog"`
	Jwt extensions.JWTConf `json:"jwt"`
	DB  extensions.EntConf `json:"db"`
}

type Configs struct {
	Dev  Config `json:"dev"`
	Prod Config `json:"prod"`
}

func (c *Configs) GetEnvConfig(env string) Config {
	switch env {
	case "dev":
		return c.Dev
	case "prod":
		return c.Prod
	default:
		return c.Dev
	}
}

func LoadConfig(env string) Config {
	curDir, err := os.Getwd()

	if err != nil {
		log.Panicf("no current folder")
	}

	index := strings.Index(curDir, "backend-go")

	root := "/app"
	if index > -1 {
		root = curDir[:index+len("backend-go")]
	}

	var configFile string

	flag.StringVar(&configFile, "f", filepath.Join(root, "config.yml"), "the config file")
	flag.Parse()

	if configFile == "" {
		log.Panicf("init config file error")
	}

	var c Configs
	conf.MustLoad(configFile, &c)
	return c.GetEnvConfig(env)
}
