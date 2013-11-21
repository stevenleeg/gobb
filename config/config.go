package config

import (
    "github.com/msbranco/goconfig"
    "log"
)

var Config *goconfig.ConfigFile

func GetConfig(path string) *goconfig.ConfigFile {
    if Config != nil {
        return Config
    }

    if path == "" {
        path = "gobb.conf"
    }

    var err error
    Config, err = goconfig.ReadConfigFile(path)
    if err != nil {
		log.Fatalln("Could not read config file", err)
    }

    return Config
}
