package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Config *SiteConfig

// 获取配置内容
func GetConfig(fileName string) {
	Config = &SiteConfig{}
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, Config)
	if err != nil {
		panic(err)
	}
}
