package configManager

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sync"
)

type Config struct {
	Port     string `yaml:"port"`
	Database struct {
		Url    string `yaml:"url"`
		DbName string `yaml:"dbname"`
	} `yaml:"database"`

	FilePathRoot string `yaml:"filePathRoot"`
	ServerMode string `yaml:"serverMode"`
}

var conf *Config
var once sync.Once

func GetConf() *Config {
	once.Do(func() {
		//read file
		pwd, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		configFile, err := ioutil.ReadFile(pwd + "/config/config.yml")

		if err != nil {
			panic(err)
		}
		conf = &Config{}
		err = yaml.Unmarshal(configFile, conf)

		if err != nil {
			panic(err)
		}
	})
	return conf
}
