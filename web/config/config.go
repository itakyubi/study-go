package config

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"study-go/web/utils"
)

type Config struct {
	AdminServer Server `yaml:"adminServer" json:"adminServer" default:"{\"port\":\":8008\"}"`
}

type Server struct {
	Port string `yaml:"port" json:"port"`
}

func LoadConfig(cfg interface{}) error {
	var configFile string
	flag.StringVar(&configFile, "c", "./web/config/config.yaml", "the configuration file")
	flag.Parse()

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
