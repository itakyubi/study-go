package common

import (
	"flag"
	"github.com/gin-contrib/cache/persistence"
	"time"
)

const (
	ConfFile        = "ConfFile"
	DefaultConfFile = "./web/config/config.yaml"
)

var cache = persistence.NewInMemoryStore(time.Minute * 10)

func init() {
	var configFile string
	flag.StringVar(&configFile, "c", "./web/config/config.yaml", "the configuration file")
	flag.Parse()
	cache.Set(ConfFile, configFile, -1)
}

func GetConfFile() string {
	res := DefaultConfFile
	cache.Get(ConfFile, &res)
	return res
}
