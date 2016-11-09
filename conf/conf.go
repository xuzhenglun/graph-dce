package conf

import (
	log "github.com/Sirupsen/logrus"
	"github.com/sakeven/go-env"
)

var conf Config

type Config struct {
	DceHost  string `json:"dce_host" env:"DCE_HOST"`
	Username string `json:"username" env:"USERNAME"`
	Password string `json:"password" env:"PASSWORD"`
}

func GetConf() Config {
	return conf
}

func ParseEnvConfig() error {
	err := env.Decode(&conf)
	if err != nil {
		return err
	}

	configCheck()
	return nil
}

func configCheck() {
	if conf.DceHost == "" {
		log.Fatal("DCE HOST can not be empty")
	}

	if conf.Username == "" {
		log.Fatal("DCE username can not be empty")
	}

	if conf.Password == "" {
		log.Fatal("DCE password can not be empty")
	}
}
