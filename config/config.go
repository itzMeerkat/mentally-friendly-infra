package config

import (
	"gitee.com/infra/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)
func LoadConfigs(path string, target interface{}) error {
	log.Infof("Loading configs from %s", path)
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(f, target)
	if err != nil {
		return err
	}

	return nil
}