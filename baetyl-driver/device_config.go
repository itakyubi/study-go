package main

import (
	"github.com/baetyl/baetyl-go/v2/http"
	"time"
)

type DeviceConfig struct {
	http.ClientConfig `yaml:",inline" json:",inline"`
	Interval          time.Duration `yaml:"interval" json:"interval" default:"20s"`
	Properties        []Property    `yaml:"properties" json:"properties"`
}

type Property struct {
	Name string `yaml:"name" json:"name"`
	Path string `yaml:"path" json:"path"`
}
