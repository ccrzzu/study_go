package config

import (
	"io/ioutil"
	"strings"
)


var Config = func () BaseConfig {
	var idcPath = "/data/app/idc/idc.ini"
	env := "dev"
	idc, err := ioutil.ReadFile(idcPath)
	if err == nil {
		env = strings.TrimSpace(string(idc))
	}
	if _, hasConfig := EnvConfigMap[env]; hasConfig {
		return EnvConfigMap[env]
	} else {
		return EnvConfigMap["dev"]
	}
}()
