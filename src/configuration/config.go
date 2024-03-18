package configuration

import (
	"os"

	dotenv "github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func NewConfig(filenames ...string) Config {
	err := dotenv.Load(filenames...)

	if err != nil {
		panic(err.Error())
	}

	return &configImpl{}
}
