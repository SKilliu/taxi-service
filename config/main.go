package config

import (
	"io/ioutil"
	"os"
	"simple-service/db"
	"sync"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Config is configuration interface
type Config interface {
	HTTP() *HTTP
	Log() *logrus.Entry
	Authentication() *Authentication
	DB() db.QInterface
}

// ConfigImpl is implementation of config interface
type ConfigImpl struct {
	sync.Mutex

	http *HTTP
	log  *logrus.Entry
	jwt  *Authentication
	db   db.QInterface
}

// New config creating
func New() Config {
	return &ConfigImpl{
		Mutex: sync.Mutex{},
	}
}

func UploadEnvironmentVariables(pathToConfigFile string) {
	ymlFile, err := os.Open(pathToConfigFile)
	if err != nil {
		panic(err)
	}

	defer ymlFile.Close()

	var variables = make(map[string]string)

	byteValue, err := ioutil.ReadAll(ymlFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(byteValue, &variables)
	if err != nil {
		panic(err)
	}

	for k, v := range variables {
		err := os.Setenv(k, v)
		if err != nil {
			panic(err)
		}
	}
}
