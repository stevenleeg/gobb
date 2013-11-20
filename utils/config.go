package utils

import (
    "github.com/msbranco/goconfig"
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
        FatalError(err, "Could not read config file")
    }

    return Config
}
