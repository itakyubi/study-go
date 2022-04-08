package utils

import (
	"github.com/pkg/errors"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func LoadYAML(path string, out interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.WithStack(err)
	}
	return UnmarshalYAML(data, out)
}

func UnmarshalYAML(in []byte, out interface{}) error {
	err := yaml.Unmarshal(in, out)
	if err != nil {
		return errors.WithStack(err)
	}
	err = validator.Validate(out)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
