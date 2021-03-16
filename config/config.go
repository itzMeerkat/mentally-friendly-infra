package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"reflect"
)
func LoadConfigFile(path string, target interface{}) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(f, target)
	if err != nil {
		panic(err)
	}
}

func LoadEnvVar(target interface{}) {
	targetType := reflect.TypeOf(target).Elem()
	fieldNum := targetType.NumField()
	for i:=0;i<fieldNum;i++ {
		fieldName := targetType.Field(i).Name
		val := os.Getenv(fieldName)
		reflect.ValueOf(target).Elem().Field(i).SetString(val)
	}
}
