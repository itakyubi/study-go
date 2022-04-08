package config

import "time"

type Config struct {
	AdminServer Server `yaml:"adminServer" json:"adminServer" default:"{\"port\":\":8008\"}"`
	Plugin      Plugin `yaml:"plugin" json:"plugin"`
}

type Server struct {
	Port         string        `yaml:"port" json:"port"`
	ReadTimeout  time.Duration `yaml:"readTimeout" json:"readTimeout" default:"30s"`
	WriteTimeout time.Duration `yaml:"writeTimeout" json:"writeTimeout" default:"30s"`
	ShutdownTime time.Duration `yaml:"shutdownTime" json:"shutdownTime" default:"3s"`
}

type Plugin struct {
	User string `yaml:"user" json:"user" default:"database"`
}
