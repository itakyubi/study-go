package context

import (
	"study-go/web/log"
)

type SystemConfig struct {
	Logger log.Config `yaml:"logger,omitempty" json:"logger,omitempty"`
}
