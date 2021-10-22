package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"study-go/web/common"
	"study-go/web/utils"
)

type Config struct {
	AdminServer Server `yaml:"adminServer" json:"adminServer" default:"{\"port\":\":8008\"}"`
	Plugin      Plugin `yaml:"plugin" json:"plugin"`
}

type Server struct {
	Port string `yaml:"port" json:"port"`
}

type Plugin struct {
	User string `yaml:"user" json:"user" default:"database"`
}

func LoadConfig(cfg interface{}) error {
	configFile := common.GetConfFile()
	if utils.FileExists(configFile) {
		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			return err
		}
		return yaml.Unmarshal(data, cfg)
	} else {
		return yaml.Unmarshal(nil, cfg)
	}
}
