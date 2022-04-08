package common

import (
	"study-go/web/utils"
)

func LoadConfig(cfg interface{}, files ...string) error {
	f := GetConfFile()
	if len(files) > 0 && len(files[0]) > 0 {
		f = files[0]
	}
	return utils.LoadYAML(f, cfg)
}
