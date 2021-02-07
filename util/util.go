package util

import jsoniter "github.com/json-iterator/go"

func Render(obj interface{}) string {
	s, err := jsoniter.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(s)
}