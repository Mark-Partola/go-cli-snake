package config

import (
	"log"
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type values struct {
	Env    string `yaml:"env" env-default:"production"`
	Server struct {
		Host string `yaml:"host" env-default:"localhost"`
		Port int    `yaml:"port" env-default:"8000"`
	} `yaml:"server"`
}

type config struct {
	values values
}

var once sync.Once
var instance config

func Setup() values {
	once.Do(func() {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("cannot resolve pwd")
		}

		var values values

		filename := pwd + string(os.PathSeparator) + "config.yaml"
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			log.Fatalf("config file not found: %s", filename)
		}

		if err := cleanenv.ReadConfig(filename, &values); err != nil {
			log.Fatalf("cannot read config: %s", err)
		}

		instance = config{
			values: values,
		}
	})

	return instance.values
}

func Get() values {
	return instance.values
}
