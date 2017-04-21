package global

import (
	log "github.com/cihub/seelog"
	"io/ioutil"
	"fmt"
	"gopkg.in/yaml.v2"
)

func Init() {
	data, err := ioutil.ReadFile("./etc/config.yaml")
	if err != nil {
		panic(fmt.Sprintf("Read config file error - %s", err))
	}

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		panic(fmt.Errorf("Parse config error - %s", err))
	}

	logger, err := log.LoggerFromConfigAsFile(Config.LogPath)
	if err != nil {
		panic(err)
	} else {
		log.ReplaceLogger(logger)
	}

	RedisClient = NewRedisClient(&Config.Redis)
}