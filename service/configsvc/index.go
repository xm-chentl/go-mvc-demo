package configsvc

import (
	"io/ioutil"
	"reflect"

	"github.com/xm-chentl/go-mvc-demo/contract"

	"gopkg.in/yaml.v2"
)

type yamlSvc struct {
	filePath string
}

func (y yamlSvc) GetStruct(res interface{}) (err error) {
	yamlFile, err := ioutil.ReadFile(y.filePath)
	if err != nil {
		return
	}

	cfg := make(map[string]interface{})
	if err = yaml.Unmarshal(yamlFile, &cfg); err != nil {
		return
	}

	key := reflect.TypeOf(res).Elem().Name()
	if content, ok := cfg[key]; ok {
		rt := reflect.TypeOf(res).Elem()
		rv := reflect.ValueOf(res).Elem()
		if reflect.TypeOf(content).Kind() == reflect.Map {
			data := content.(map[interface{}]interface{})
			for i := 0; i < rt.NumField(); i++ {
				if v, ok := data[rt.Field(i).Name]; ok {
					rv.Field(i).Set(
						reflect.ValueOf(v),
					)
				}
			}
		}
	}

	return
}

func NewYaml(src string) contract.IConfig {
	return &yamlSvc{
		filePath: src,
	}
}
