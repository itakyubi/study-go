package common

import (
	"github.com/gin-contrib/cache/persistence"
	"time"
)

const (
	ConfFile        = "ConfFile"
	DefaultConfFile = "./web/etc/conf.yml"
)

var cache = persistence.NewInMemoryStore(time.Minute * 10)

func SetConfFile(v string) {
	cache.Set(ConfFile, v, -1)
}

func GetConfFile() string {
	res := DefaultConfFile
	cache.Get(ConfFile, &res)
	return res
}
