package config

import (
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)
func LoadConfigs(log *zap.SugaredLogger, path string, target interface{}) error {
	log.Infof("Loading configs from %s", path)
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	//target = reflect.ValueOf(target).Elem()
	err = yaml.Unmarshal(f, target)
	if err != nil {
		return err
	}

	return nil
}